package main

import "fmt"

func main() {
	const name = "pynickle"
	const sex string = "man"
	const (
		a = iota
		b = iota
	)
	fmt.Println(a, b)
	const (
		c = iota
		d = "string"
		e
		f = iota
	)
	fmt.Println(c, d, e, f)
}
