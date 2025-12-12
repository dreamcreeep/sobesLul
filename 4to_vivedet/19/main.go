package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int    `json:"one"`
	two string `json:"two"`
}

func testData() {
	in := MyData{1, "two"}

	fmt.Printf("%#v\n", in) // MyData{1, "two"}

	encoded, _ := json.Marshal(in)

	fmt.Println(string(encoded)) // One:1

	var out MyData

	json.Unmarshal(encoded, &out)

	fmt.Printf("%#v\n", out)
}

func main() {
	testData()
}
