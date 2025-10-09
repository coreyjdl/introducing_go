package main

import "fmt"

func main() {

	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())


	fmt.Println("Next Even Numbers")
	next := nextEven()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}



func nextEven() func() int {
	x := 0
	return func() int {
		x += 2
		return x
	}
}