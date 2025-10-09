package main

import "fmt"
import "time"

func main() {

	// Create two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Start two goroutines that send messages to the channels at different intervals
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	// Start a goroutine that uses select to wait on both channels
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	// Prevent program from exiting
	var input string
	fmt.Scanln(&input) // wait for Enter Key
}