package main

import (
	"fmt"
)

func main() {
	var countryCapitalMap map[string]string

	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Captital of", country, "is", countryCapitalMap[country])
	}

	captial, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Captital of United States is", captial)
	} else {
		fmt.Println("Captital of United States is not present")
	}
}
