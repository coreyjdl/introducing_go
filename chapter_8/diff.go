package main

import ("fmt"; "hash/crc32"; "io"; "os")

func getHash(filePath string) (uint32, error) {
	
	// Read the file content
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	
	// ensure the file is closed after reading
	defer f.Close()

	// create the hasher
	hasher := crc32.NewIEEE()

	// copy data
	if _, err := io.Copy(hasher, f); err != nil {
		return 0, err
	}

	// return the checksum
	return hasher.Sum32(), nil
}

func main() {
	h1, err := getHash("file1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	h2, err := getHash("file2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if h1 == h2 {
		fmt.Println("Files are identical")
	} else {
		fmt.Println("Files differ")
	}
}