package main
import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("write.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	file.WriteString("Hello, Go file handling!\n")
	fmt.Println("File written successfully.")
}