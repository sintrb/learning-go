package main

import (
	"fmt"
)

func go1(ch chan int) {
	println("go1")
	ch <- 4
}

func go2(ch chan int) {
	println("go2")
	ch <- 2
}

func main() {
	println("main")
	ch := make(chan int)
	go go1(ch)
	go go2(ch)
	fmt.Println(fmt.Sprintf("--%d--", <-ch))
	fmt.Println(fmt.Sprintf("--%d--", <-ch))
}
