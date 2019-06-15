package main

import "fmt"

func swap(x *int, y *int) {
	var temp int
	temp = *x  
	*x = *y     
	*y = temp
}

func main() {
	var a int = 1
	fmt.Println(&a)
	var ip *int
	fmt.Println(ip, ip==nil)
	ip = &a
	fmt.Println(ip, *ip)
	
	var array = []int{10,20,30}
	var apointer [3]*int
	var i int
	for i=0;i<3;i++{
		apointer[i] = &array[i]
	}
	for i=0;i<3;i++{
		fmt.Println(*apointer[i])
	}
	
	var pointer1 *int
	var pointer2 **int
	pointer1 = &a
	pointer2 = &pointer1
	fmt.Println(*pointer1, **pointer2)
	
	first := 10
	second := 20
	swap(&first, &second)
	fmt.Println(first, second)
}
