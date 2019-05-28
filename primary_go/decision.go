package main

import ("fmt"
	"time")

func main(){
	var a int = 10
	if a<0{
		fmt.Println("a is under zero")
	}else if a==0{
		fmt.Println("a is zero")
	}else{
		fmt.Println("a is above zero")
	}
	switch{
		case a<10:
			fmt.Println("a is smaller than 10")
		case a>10:
			fmt.Println("a is bigger than 10")
		default:
			fmt.Println("invalid or it's exactly 10")
	}
    //below code from https://www.jianshu.com/p/2a1146dc42c3
	start := time.Now()
    c := make(chan interface{})
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {

        time.Sleep(4*time.Second)
        close(c)
    }()

    go func() {

        time.Sleep(3*time.Second)
        ch1 <- 3
    }()

      go func() {

        time.Sleep(3*time.Second)
        ch2 <- 5
    }()

    fmt.Println("Blocking on read...")
    select {
		case <- c:
		    fmt.Println("Unblocked %v later.\n", time.Since(start))
		case <- ch1:
		    fmt.Println("ch1 case...")
		case <- ch2:
		    fmt.Println("ch2 case...")
		default:
		    fmt.Println("default go...")
    }
    start = time.Now()
    c = make(chan interface{})
    ch1 = make(chan int)
    ch2 = make(chan int)

    go func() {

        time.Sleep(4*time.Second)
        close(c)
    }()

    go func() {

        time.Sleep(3*time.Second)
        ch1 <- 3
    }()

    go func() {

        time.Sleep(3*time.Second)
        ch2 <- 5
    }()

    fmt.Println("Blocking on read...")
    select {
    case <- c:

        fmt.Println("Unblocked %v later.\n", time.Since(start))

    case <- ch1:

        fmt.Println("ch1 case...")
      case <- ch2:

        fmt.Println("ch2 case...")
    //default:

        //fmt.Println("default go...")
    }
}
