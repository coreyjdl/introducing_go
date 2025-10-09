package main

import ("fmt"; "crypto/sha256")

func main() {
	// create a new hash.Hash computing the SHA256 checksum
	h := sha256.New()

	// write a byte slice to it
	h.Write([]byte("mySuperSecretPassword"))

	// get the resulting checksum
	bs := h.Sum([]byte{})

	// print it out
	fmt.Println(bs)

	// get input from the user
	var input string
	fmt.Println("Enter text to hash:")
	fmt.Scanln(&input)
	
	// check it against the original
	h2 := sha256.New()
	h2.Write([]byte(input))
	bs2 := h2.Sum([]byte{})

	if string(bs) == string(bs2) {
		fmt.Println("You guessed it!")
	} else {
		fmt.Println("Try again.")
	}	
}