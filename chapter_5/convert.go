package main
import "fmt"

func main() {

	x := [5]float64 {
		90, 99, 55, 88, 75,
	}

	// x[0] = 90
	// x[1] = 99
	// x[2] = 55
	// x[3] = 88
	// x[4] = 75


	var total float64 = 0

	for i:=0; i< len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))

}
