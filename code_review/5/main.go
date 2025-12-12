package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Необходимо хранить список из username и почты в памяти программы.
// Пользователи программы могут добавлять и просматривать userов.
// Разработчик написал данный код, но не прошел ревью.
// Нужно указать на ошибки и исправить, либо оставить комментарий о том, что можно улучшить.

var users = []string{}

type User struct {
    Username string json:"username"
    email    string json:"email"
}

func addUser(w http.ResponseWriter, r *http.Request) {
    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Error parsing JSON", http.StatusBadRequest)
        return
    }

    users = append(users, newUser.Username)

    fmt.Fprintf(w, "Added user")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/addUser", addUser)
    http.HandleFunc("/getUsers", getUsers)
    fmt.Println("Server starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
