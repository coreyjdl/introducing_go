package main

import "fmt"

var x string = "test"
const y string = "hello"

func main() {
  fmt.Println(x)  
  fmt.Println(y)
  f()
}

func f() {
	fmt.Println(x)
}
