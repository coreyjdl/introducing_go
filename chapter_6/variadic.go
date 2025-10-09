package main
import "fmt"


func main() {
	result := add(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}