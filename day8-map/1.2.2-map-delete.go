package main

import (
	"fmt"
)

func main() {
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "China": "Beijing"}
	fmt.Println("origin map")
	for country := range countryCapitalMap {
		fmt.Println("Captital of", country, "is", countryCapitalMap[country])
	}

	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("new map")
	for country := range countryCapitalMap {
		fmt.Println("Captital of", country, "is", countryCapitalMap[country])
	}
}
