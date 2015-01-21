package main

import "time"

func source(strs []string) <-chan string {
	out := make(chan string, 10000)
	go func() {
		for _, s := range strs {
			out <- s
			println("<-", s)
			time.Sleep(1000)
		}
		println("close")
		close(out)
	}()
	println(out)
	println(&out)
	return out
}

func main() {
	o := source([]string{"Hello", "World"})
	time.Sleep(1000)
	println(<-o)
	time.Sleep(1000)
	println(<-o)
	time.Sleep(1000)
	println(<-o)
}
