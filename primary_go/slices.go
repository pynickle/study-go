package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[:2])
	fmt.Println(s[1:4])
	fmt.Println(s[3:])
	fmt.Println(len(s), cap(s))
	s2 := make([]int, len(s), (cap(s))*2)
	copy(s2, s)
	fmt.Println(s2, len(s2), cap(s2))
}
