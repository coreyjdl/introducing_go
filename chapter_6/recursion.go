package main

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(0))
	fmt.Println(factorial(1))
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 5! = 5 * 4 * 3 * 2 * 1
// 0! = 1
// 1! = 1
// 10! = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1