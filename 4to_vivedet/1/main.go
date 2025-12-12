package main

import "fmt"

func screen(nums []int, pos, num int) []int {
	pref := nums[:pos] // [1,2] len2, cap5

	post := nums[pos:] // [3,4,5] len3, cap3
	fmt.Println(len(post), cap(post))

	pref = append(pref, num) //[1,2,10]

	return append(pref, post...) // [1,2,10,10,4,5]
}

func main() {
	fmt.Println(screen([]int{1, 2, 3, 4, 5}, 2, 10))
}
