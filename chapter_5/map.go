package main

import "fmt"

func main() {
	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

}