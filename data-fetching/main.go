package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	launchesEndpoint := "https://api.spacexdata.com/v3/launches"

	res, err := http.Get(launchesEndpoint)
	if err != nil {
		log.Fatalf("Error making GET request: %s", err)
	}
	// Prevent resource leaks
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		bodyString := string(bodyBytes)
		fmt.Printf("Response body: %v", bodyString)
	} else {
		fmt.Printf("Received non-OK status code: %v\n", res.StatusCode)
	}
}