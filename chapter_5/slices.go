package main

import "fmt"

func main() {
	arr := []float64 {
		90, 99, 55, 88, 75,
	}

	x := arr[0:3]
	fmt.Println(x)
	
	y := arr[1:4]
	fmt.Println(y)

	z := arr[:5]
	fmt.Println(z)
}