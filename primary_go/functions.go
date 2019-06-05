package main

import "fmt"

func max(a, b int) int {
	var result int
	if a > b {
		result = a
	} else if a == b {
		result = a + b
	} else {
		result = b
	}
	return result
}

func reverse(x, y string) (string, string) {
	return y, x
}

func main() {
	result1 := max(1, 5)
	fmt.Println(result1)
	first, second := reverse("hello", "world")
	fmt.Println(first, second)
}
