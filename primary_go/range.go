package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    kvs := map[string]string{"name": "nick", "age": "13"}
    for k, v := range kvs {
        fmt.Printf("%s : %s\n", k, v)
    }
}