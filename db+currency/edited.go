package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

type Bank struct {
	Name    string
	CurFrom string
	CurTo   string
	URL     string
	Auth    bool
}

func main() {
	if len(os.Args) != 2 || os.Args[1] != "update" {
		fmt.Println("Usage: ./currency update")
		return
	}

	cfg := Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     5432,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	defer db.Close()

	banks := []Bank{
		{"Bank 1", "RUB", "USD", "http://bank.example.com/rates/rub-usd", false},
		{"Bank 2", "RUB", "USD", "http://bank2.example.com/rates?currFrom=RUR&currTo=USD", true},
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Graceful shutdown...")
			return
		case <-ticker.C:
			runUpdate(ctx, db, client, banks)
		}
	}
}

func connectDB(cfg Config) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to DB")
	return db, nil
}

func runUpdate(ctx context.Context, db *sql.DB, client *http.Client, banks []Bank) {
	var wg sync.WaitGroup

	for _, bank := range banks {
		wg.Add(1)
		go func(b Bank) {
			defer wg.Done()
			if err := fetchAndSave(ctx, db, client, b); err != nil {
				log.Printf("failed to update %s: %v", b.Name, err)
			}
		}(bank)
	}
	wg.Wait()
}

func fetchAndSave(ctx context.Context, db *sql.DB, client *http.Client, bank Bank) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bank.URL, nil)
	if err != nil {
		return err
	}
	if bank.Auth {
		req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	strBody := string(body)
	if bank.Name == "Bank 1" {
		strBody = strings.ReplaceAll(strBody, ",", ".")
	}

	value, err := strconv.ParseFloat(strBody, 64)
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO currency_rates (bank, "from", "to", value) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, bank.Name, bank.CurFrom, bank.CurTo, value); err != nil {
		return err
	}

	return tx.Commit()
}
