package main

import "fmt"

func location(city string) (string, string) {
	var region string
	var weather string

	switch city {
	case "LA", "SF", "SD":
		region, weather = "Westcoast", "Sun and Beer!!"
	case "BOS", "NYT", "CHI":
		region, weather = "Eastcoast", "Crowded city and heavy Snow! Happy!"
	default:
		region, weather = "I don't know the place", "the weather I don't know."
	}
	return region, weather
}

func main() {
	region, weather := location("MARS")
	fmt.Printf("It's in %s. That has %s", region, weather)
}
