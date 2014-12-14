package main

// import "fmt"

func printarr(arr [10]interface) {

}

func main() {
	var v float32
	v = 2000.2
	println(v)

	var c complex64
	c = 3.2 + 84i
	println(c)

	str := "中国Hello World"
	println(str)

	s := str[0:3]
	println(s)

	arr := []int{1, 3, 5}
	for _, v := range arr {
		println(v)
	}

	println("------")
	arr = append(arr, 3)
	for _, v := range arr {
		println(v)
	}
}
