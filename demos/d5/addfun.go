package main

// import "fmt"

type Interger int

func (a *Interger) add(b Interger) {
	*a += b
}

func main() {
	var v Interger = 4

	println(v)

	v.add(6)

	println(v)
}
