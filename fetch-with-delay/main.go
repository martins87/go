package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// url := "https://www.google.com"
	url := "https://httpbin.org/status/300"

	res, err := httpCallWithDelay(url)
	if err != nil {
		fmt.Printf("Error on the request: %v\n", err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("Request status code: %v\n", res.StatusCode)
}

func httpCallWithDelay(url string) (*http.Response, error) {
	const trials = 3
	const delay = 2 * time.Second

	for range trials {
		res, err := http.Get(url)
		if err == nil && res.StatusCode >= 200 && res.StatusCode < 300 {
			return res, err
		}

		if err != nil {
			fmt.Printf("Request error: %v", err)
		} else {
			fmt.Printf("Unsuccessful request: %v\n", res.StatusCode)
			res.Body.Close()
		}

		time.Sleep(delay)
	}

	return nil, fmt.Errorf("Failed after %d attempts", trials)
}