package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3} // 3,3
	s2 := s1[1:]         // [2,3] 2,2
	s2 = append(s2, 4)   //[2,3,4] 3,4
	s2[0] = 10           // [10,3,4]

	fmt.Println(s1) // 1,2,3
	fmt.Println(s2) // 10,3,4
}
