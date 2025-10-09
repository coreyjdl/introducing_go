package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HomePageSize struct to hold URL and its size
type HomePageSize struct {
	URL string
	Size int
}

func main() {

	// List of URLs to check
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.reddit.com",
		"https://www.amazon.com",
	}

	// Create a buffered channel with capacity equal to number of URLs
	results := make(chan HomePageSize)

	// Start a goroutine for each URL to fetch its content
	for _, url :=range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{url, len(bs)}
		}(url)
	}


	// Collect results and determine the biggest homepage
	var biggest HomePageSize
	for range urls {
		result := <- results
		if result.Size > biggest.Size {
			biggest = result
		}	
	}

	// Print the biggest homepage
	fmt.Printf("The biggest homepage is %s with a size of %d\n", biggest.URL, biggest.Size)
}