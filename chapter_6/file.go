package main
import (
	"os"
	"io"
	"fmt"
)


func main() {

	f, _ := os.Open("ex.csv")
	defer f.Close()

	content, _ := io.ReadAll(f)
	fmt.Println(string(content))

}