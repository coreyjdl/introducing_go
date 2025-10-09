package main

import "fmt"

func main() {

	i, err := fmt.Println("Hello, World")

	fmt.Println("i:", i)
	fmt.Println("err:", err)


	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(3, 4))

}

