package main

import "fmt"

func main() {
	var i int
	fmt.Scanln(&i)
	switch {
	case i < 0:
		fmt.Println("<0")
	case i == 0:
		fmt.Println("==0")
	case i > 0:
		fmt.Println(">0")
	}

}
