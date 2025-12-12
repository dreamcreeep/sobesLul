package main

/*
Есть заказы, которые в процессе жизненного цикла переходят по статусам.

Требуется реализовать слой работы с БД:
- Сохранять историю статусов заказа
- Учесть партицирование

В дальнейшем, за рамками текущей задачи, будут реализованы другие методы:
получение статусов по времени, полная история заказа и т.п.
*/

import (
	"database/sql"
	"fmt"
	"time"
)

type OrderStatusHistoryStore struct {
	db *sql.DB
}

func NewStore(db sql.DB) (OrderStatusHistoryStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ping failed: %w", err)
	}

	return &OrderStatusHistoryStore{db: db}, nil
}

func (s *OrderStatusHistoryStore) SaveOrderStatusHistory(orderId, status) {
	Query := "insert into orderstatushistory (orderid, status, insertedat) values($1, $2, $3, $4)"

	_, err := s.db.Exec(query, orderid, status, time.Now())

	if err != nil {
		return fmt.Errorf("failed save order: %w", err)
	}
	return nil
}
