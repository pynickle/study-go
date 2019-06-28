package main

import "fmt"

func main() {
    var info map[string]string
    info = make(map[string]string)

    info [ "nick" ] = "13"
    info [ "alice" ] = "10"
    info [ "jack" ] = "16"
    info [ "john" ] = "19"

    for name := range info {
        fmt.Println(name + "'s age is " + info [name])
    }

    delete(info, "jack")
    fmt.Printf("\n")
    for name := range info {
        fmt.Println(name + "'s age is " + info [name])
    }
}