package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Rocket struct {
	RocketId 			string			`json:"rocket_id"`
	RocketName 		string	`json:"rocket_name"`
}

type Launch struct {
	FlightNumber	int			`json:"flight_number"`
	MissionName		string	`json:"mission_name"`
	LaunchYear		string	`json:"launch_year"`
	LaunchSuccess	bool		`json:"launch_success"`
	Details				string	`json:"details"`
	Rocket								`json:"rocket"`
}

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

		var launches []Launch
		json.Unmarshal(bodyBytes, &launches)

		for _, launch := range launches {
			fmt.Printf("%+v\n\n", launch)
		}
	} else {
		fmt.Printf("Received non-OK status code: %v\n", res.StatusCode)
	}
}
