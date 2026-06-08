package main

import (
	"fmt"
	"log"
	"os"

	"github.com/martins87/signature-verifier/internal/bitcoin"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		log.Fatalf("usage: %s <address> <message> <signature>", args[0])
	}

	req := bitcoin.VerifyRequest{
		Address:   args[1],
		Message:   args[2],
		Signature: args[3],
	}

	valid, err := bitcoin.Verify(req)
	if err != nil {
		log.Fatalf("error verifying signature: %s", err)
	}

	if valid {
		fmt.Println("Signature is VALID")
	} else {
		fmt.Println("Signature is NOT VALID")
	}
}
