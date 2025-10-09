package main
import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // make a slice with length 2
	copy(slice2, slice1) // copy from slice1 to slice2
	fmt.Println(slice2)
}