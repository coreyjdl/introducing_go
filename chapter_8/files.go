package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := "file.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// read file size
	bs := make([]byte, fileInfo.Size())
	_, err = file.Read(bs)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}