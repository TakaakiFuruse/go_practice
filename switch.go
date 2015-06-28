package main

import (
	"fmt"
)

var (
	region  string
	weather string
	city_ar = []string{"LAX", "SFC", "SAN", "ORD",
		"NRT", "ROM", "SAN", "NYC"}
)

func location(city string) (string, string) {
	switch city {
	case "LAX", "SFC", "SAN":
		region, weather = "Westcoast", "Sun and Beer!!"
	case "BOS", "NYC", "ORD":
		region, weather = "Eastcoast", "Crowded city and heavy Snow! Happy!"
	default:
		region, weather = "the place where I don't know.", "the weather I don't know."
	}
	return region, weather
}

func return_answer(pLAXce_name string) {
	region, weather := location(pLAXce_name)
	fmt.Printf("%s? It's in %s. That has %s \n", pLAXce_name, region, weather)
}

func main() {
	for i := 0; i < len(city_ar); i++ {
		return_answer(city_ar[i])
	}
}
