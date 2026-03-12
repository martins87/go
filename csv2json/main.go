package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Record struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func main() {
	csvFilepath := "data.csv"

	// Read csv file
	csvFile, err := os.Open(csvFilepath)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", csvFilepath, err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var records []Record
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading line: %v", err)
		}

		record := Record{
			Name: row[0],
			Email: row[1],
			PhoneNumber: row[2],
		}
		records = append(records, record)
	}

	fmt.Printf("Records: %+v", records[1:len(records)-1])

	// Marshal struct to JSON bytes
	data, err := json.Marshal(records)
	if err != nil {
		log.Fatalf("Error serializing records: %v", err)
	}

	// Convert bytes to string
	jsonString := string(data)
	fmt.Printf("records as json: %v\n", jsonString)

	jsonFilepath := "data.json"
	err = os.WriteFile(jsonFilepath, data, 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", jsonFilepath, err)
	}
	fmt.Printf("Data successfully writen to %s\n", jsonFilepath)
}
