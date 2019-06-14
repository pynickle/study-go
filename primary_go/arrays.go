package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var arrays1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(arrays1)
	var arrays2 = [...]string{"Hello", "World"}
	fmt.Println(arrays2)
	var i,j int
	for i=0;i<5;i++{
		fmt.Println(i, arrays1[i])
	}
	var arrays3 = [2][5]int{
		{1,2,3,4,5},
		{6,7,8,9,0},
	}
	fmt.Println(arrays3)
	for  i = 0; i < 2; i++ {
		for j = 0; j < 5; j++ {
			fmt.Println(i, j, ":", arrays3[i][j])
		}
	}
}
