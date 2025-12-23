// Нужно оставить только уникальные!

package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
)

type Order struct {
 ID     string json:"id"   // ID заказа (поле в JSON)
 Amount int    json:"amount" // Сумма заказа (поле в JSON)
}

func loadOrders(filename string) ([]Order, error) {
 // Читаем содержимое файла
 data, err := ioutil.ReadFile(filename)
 if err != nil {
  return nil, err
 }

 var orders []Order
 // Десериализуем JSON в массив структур Order
 err = json.Unmarshal(data, &orders)
 if err != nil {
  return nil, err
 }

 return orders, nil
}


func processOrders(orders []Order) {
 for _, order := range orders {
  fmt.Println("Processing order:", order.ID)
 }
}

func main() {
 // Загружаем заказы из файла orders.json
 orders, err := loadOrders("orders.json")
 if err != nil {
  fmt.Println("Error loading orders:", err)
  return
 }

 // Обрабатываем заказы (выводим их в консоль)
 processOrders(uniqueOrders)
}
