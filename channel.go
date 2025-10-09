package main

import "fmt"

func sum(c chan int) {
	x := <- c
	y := <- c
	c <- x + y
}

func main() {
	c:= make(chan int)
	go sum(c)
	c <- 10
	c <- 15
	r := <- c

	fmt.Println(r)
}
