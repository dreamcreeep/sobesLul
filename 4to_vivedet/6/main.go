package main

import "fmt"

type errorString struct {
	code int
}

func (e errorString) Error() string {
	return fmt.Sprintf("error message: %v", e.code)
}

func checkErr(err error) bool {
	return err != nil
}

func main() {
	var e *errorString
	fmt.Printf("Val %v, Type %T", e, e)

	if checkErr(e) {
		fmt.Println(1, e)
	}

	e = new(errorString)

	fmt.Printf("Val %v, Type %T", e, e)
	if checkErr(e) {
		fmt.Println(2, e)
	}

	e = nil
	fmt.Printf("Val %v, Type %T", e, e)
	if !checkErr(e) {
		fmt.Println(3, e)
	}
}
