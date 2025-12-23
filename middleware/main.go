package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "time"
)

func main() {
    timeout := 1 * time.Second
    router := http.NewServeMux()
    router.HandleFunc("/orders/", rateLimit(listOrders))
    srv := httptest.NewServer(router)

    for i := 0; ; i++ {
        res, err := http.Get(fmt.Sprintf("%s/orders?id=%d", srv.URL, i))
        if err != nil {
            panic(err)
        }

        if res.StatusCode == http.StatusOK {
            p, _ := io.ReadAll(res.Body)
            res.Body.Close()
            fmt.Println(string(p))
        }
    }
}

// Реализовать мидлвару
// Редактировать можно только этот метод
// В результате при обращении к данной ручке ответ должен приходить по таймауту
// В остальных случаях возвращать ошибку, например 429

func rateLimit(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        handler(w, r)
    }
}

func listOrders(w http.ResponseWriter, r *http.Request) {
    var order struct {
        ID string json:"id"
    }

    id := r.URL.Query().Get("id")
    order.ID = id
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&order)
}
