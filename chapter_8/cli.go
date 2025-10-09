package main

import ("fmt"; "flag"; "math/rand")

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "maximum number")
	flag.Parse()

	fmt.Println(rand.Intn(*maxp))
}