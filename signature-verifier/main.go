package main

import (
	"fmt"
	"log"
	"os"
)

const filename = "message.txt"

func main() {
	args := os.Args[1:]
	address := args[0]
	signature := args[1]

	fmt.Printf("Address: %s\n", address)
	fmt.Printf("Signature: %s\n", signature)

	message, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s", err)
	}

	fmt.Printf("Message: %s", string(message))
}
