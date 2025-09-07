package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// ТЗ: мы хотим собирать информацию по курсам валют в разных банках.
// Требуется написать программу, которая каждую минуту будет отправлять запрос в банк и получать курс нескольких валют и сохранять результат в БД.
// Банков может быть несколько.

func main() {

	if len(os.Args) == 2 { // желательно корректно завершать работу программу с помощью graceful shutdowp
		cmdName := os.Args[1]
		if cmdName == "help" {
			fmt.Println("Usage is './currency update'")

		} else if cmdName == "update" {

			urlsBank := []struct {
				bankName string
				curFrom  string
				curTo    string
				url      string
			}{
				{
					bankName: "Bank 1",
					curFrom:  "RUB",
					curTo:    "USD",
					url:      "http://bank.example.com/rates/rub-usd",
				},
				{
					bankName: "Bank 2",
					curFrom:  "RUB",
					curTo:    "USD",
					url:      "http://bank2.example.com/rates?currFrom=RUR&currTo=USD",
				},
			}

			clientBank := &http.Client{
				Timeout: 10 * time.Minute,
			}

			for _, url := range urlsBank { // нужно ассинхронно делать запросы с помощью горутин

				req, err := http.NewRequest(http.MethodGet, url.url, nil)
				if err != nil {

				}

				if url.bankName == "Bank 2" {
					req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
				}

				resp, err := clientBank.Do(req)

				if err != nil {
					panic(err) // не хорошо что наше приложение может в любой момент упасть
				}

				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					panic(err) // не хорошо что наше приложение может в любой момент упасть
				}

				strBody := string(body)

				if url.bankName == "Bank 1" {
					strBody = strings.ReplaceAll(strBody, ",", ".")
				}

				value, err := strconv.ParseFloat(strBody, 64)
				if err != nil {
					panic(err) // не хорошо что наше приложение может в любой момент упасть
				}
				err = updateCurrency(url.bankName, url.curFrom, url.curFrom, value)
				if err != nil {
					panic(err) // не хорошо что наше приложение может в любой момент упасть
				}
			}
		}
	} else {
		fmt.Println("Usage is './currency update'")
	}
}

const ( // вынести в конфиг
	host     = "localhost" // нельзя передавать в прод локалхост, должны получать из внешнего хранилица
	port     = 5432
	user     = "postgres"
	password = "<password>"
	dbname   = "<dbname>"
)

func updateCurrency(bank, from, to string, value float64) error { // выполнять в транзакции
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn) // конекшин в начале программы а не каждый раз
	CheckError(err)

	defer db.Close() // тоже самое нельзя закрывать соединение каждый раз

	err = db.Ping() // делать пинг сразу в начале программы и не делать при каждом вызове пинг
	CheckError(err)

	fmt.Println("Connected!") // логировать отдельной библиотекой

	insertStmt := fmt.Sprintf(`insert into currency_rates ("bank", "from", "to", "value") values('%s', '%s', '%s', '%.2f')`, bank, from, to, value) // экранировать избежать sql инъекции
	_, err = db.Exec(insertStmt)
	return err
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
