package main

import "fmt"

var x string

func main() {
  
	for i:=1; i<=10; i+=1 {
		switch i {
		case 0: x = "zero"
		case 1: x = "one"
		case 2: x = "two"
		case 3: x = "three"
		case 4: x = "four"
		case 5: x = "five"
		case 6: x = "six"
		case 7: x = "seven"
		case 8: x = "eight"
		case 9: x = "nine"
		case 10: x= "ten"
		default: x = "unknown"
	}
	fmt.Println(x)

	}
}
