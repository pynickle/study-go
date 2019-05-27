package main

import "fmt"

func more(a,b,c int) (int,int,int,int) {
	var result int = a + b + c
	return a,b,c,result
}

func main() {
	var a string = "string"
	fmt.Println(a)
	var b,c int = 5,7
	fmt.Println(b,c)
	var d = true
	fmt.Println(d)
	var e bool
	fmt.Println(e)
	f := "use :="
	fmt.Println(f)
	var(
		g int
		h string
	)
	fmt.Println(g,h)
	_,x,y,z := more(1,2,3)
	fmt.Println(x,y,z)
}
