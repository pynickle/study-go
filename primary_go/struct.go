package main

import "fmt"

type Person struct{
	age int
	name string
	job string
}

func main(){
	Nick1 := Person{13, "nick", "student"}
	fmt.Println(Nick1)
	Nick2 := Person{age:13, name:"nick", job:"student"}
	fmt.Println(Nick2)
	var Nick3 Person
	Nick3.age = 13
	Nick3.name = "nick"
	Nick3.job = "student"
	fmt.Println(Nick3)
	printperson(&Nick3)
}

func printperson(ps *Person){
	fmt.Println(ps.age)
	fmt.Println(ps.name)
	fmt.Println(ps.job)
}