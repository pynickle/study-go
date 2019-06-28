package main

import "fmt"

func pow(x uint64)(result uint64) {
    if (x > 0) {
        result = x * pow(x-1)
        return result
    }
    return 1
}

func main() {  
    var i uint64 = 5
    fmt.Printf("pow(%d) is %d\n", i, pow(i))
}