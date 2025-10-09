package main

import "fmt"

func main() {
	xs := []int{98,93,77,82,83}

	r, s := average(xs)
	 fmt.Println(r, s)	
}


func average(xs []int) (r float64, s float64) {
	total := 0.0
	for _, v := range xs {
		total += float64(v)
	}
	s = total
	r = total / float64(len(xs))

	return
}