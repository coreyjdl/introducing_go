package main

import "fmt"

func main() {

  var x[5] float64
  x[0] = 98
  x[1] = 93
  x[2] = 89
  x[3] = 77
  x[4] = 55

  var total float64 = 0
  for i:=0; i<5;  i+=1 {
	total += x[i]
	}

  fmt.Println(total /5)
}

