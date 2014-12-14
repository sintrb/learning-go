package main

type Foo struct {
	name string
	age  int
}

func main() {
	u := new(Foo)
	u.name = "Robin"
	u.age = 24
	println(u.name)
	println(u.age)
}
