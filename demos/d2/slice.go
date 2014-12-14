package main

func main() {
	var arr []int = []int{1, 2, 3, 4}
	arr = append(arr, 1)
	println(arr)

	arr2 := make([]int, 2)
	arr2[0] = 3
	arr2[0] = 2
	arr2 = append(arr2, 3)
	println(arr2)
}
