package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name     string
	Surname  string
	address  string   `json:"address"`
	Phone    string   `json:"-"`
	Parents  []string `json:"parents"`
	Children []string `json:"children"`
	Extra    string   `json:"extra"`
}

func main() {
	p := Person{
		name:     "John",
		Surname:  "Doe",
		address:  "Moscow",
		Phone:    "+8123123",
		Children: make([]string, 0),
	}

	res, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}
