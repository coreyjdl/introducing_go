package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}	
	slice1 = append(slice1, slice2...) // unpacking	
	fmt.Println(slice1)


	slice4 := append(slice1, 7, 8, 9)
	fmt.Println(slice4)
}