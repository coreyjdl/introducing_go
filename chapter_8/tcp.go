package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

func server() {
	// listen on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
 
	defer ln.Close()

	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// handle connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// recieve data
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		c.Close()
		return
	}
	fmt.Println("Received:", msg)
	c.Close()
}

func client() {
	// connect to server
	c, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer c.Close()

	// send data
	msg := "Hello, Server!"
	fmt.Println("Sending:", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}
	fmt.Println("Message sent")
}

func main() {

	go server()
	go client()

	var input string
	fmt.Scanln(&input)

}