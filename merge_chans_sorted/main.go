//Напишите функцию mergeSorted, которая принимает на вход два "отсортированных" канала и возвращает результирующий
//отсортированный канал.

package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {

}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)
	for val := range c {
		fmt.Printf("%d ", val)
	}
}

//Вывод: -1 1 2 4 4 5
