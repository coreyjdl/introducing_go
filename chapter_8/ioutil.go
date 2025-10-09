package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
