package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Rocket struct {
	RocketId string `json:"rocket_id"`
	RocketName string `json:"rocket_name"`
}

type Launch struct {
	FlightNumber int `json:"flight_number"`
	MissionName string `json:"mission_name"`
	LaunchYear string `json:"launch_year"`
	LaunchSuccess bool `json:"launch_success"`
	Details string `json:"details"`
	Rocket `json:"rocket"`
}

type LaunchList []Launch

func main() {
	year := flag.String("y", "2006", "config string")
	flag.Parse()

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

		var launches LaunchList
		json.Unmarshal(bodyBytes, &launches)

		launchesByYear := launches.getLaunchesByYear(*year)

		printLaunches(launchesByYear)
	} else {
		fmt.Printf("Received non-OK status code: %v\n", res.StatusCode)
	}
}

func printLaunches(l LaunchList) {
	for _, launch := range l {
		fmt.Printf("%+v\n\n", launch)
	}
}

func (l LaunchList) getLaunchesByYear(y string) []Launch {
	list := []Launch{}

	for _, launch := range l {
		if launch.LaunchYear == y {
			list = append(list, launch)
		}
	}

	return list
}
