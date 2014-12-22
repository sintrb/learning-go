package main

import "fmt"

func Add(a, b int) (r int) {
	r = a + b
	return r
}

func TwoRet() (a, b int) {
	a = 2
	b = 5
	return a, b
}

func manyargs(args ...interface{}) {
	for ix, arg := range args {
		fmt.Printf("%d->%v\n", ix, arg)
	}
}

func genFunc() func() {
	f := func() {

	}
	return f
}

func main() {
	manyargs(1, 2, 3)
	manyargs("I", "LOVE", "U")

	println(Add(1, 4))
	a, b := TwoRet()
	println(a)
	println(b)

	s := genFunc()
	println(s)
	s()
}
