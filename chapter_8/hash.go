package main
import ("fmt"; "hash/crc32")

func main() {
	// create a hasher
	h := crc32.NewIEEE()

	// write our data to it
	h.Write([]byte("test"))

	// calculate the checksum
	v:= h.Sum32()

	// print it
	fmt.Printf("CRC32: %d\n", v)

}