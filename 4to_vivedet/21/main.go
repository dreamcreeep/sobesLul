package main

import "fmt"

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate() error {
	var err CustomError

	if c.Age < 0 {
		err = CustomError{msg: "Unvalid age"}
	}

	if c.Name == "" {
		err = CustomError{msg: "Empty name"}
	}

	return err
}

func main() {
	c := Customer{
		Age:  1,
		Name: "Test",
	}

	err := c.Validate()
	if err != nil {
		fmt.Println("Unvalid user data:", err.Error())
	}
}
