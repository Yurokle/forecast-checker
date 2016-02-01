package main

import (
	"fmt"
	"forecast-checker/fetcher"
	"forecast-checker/conditions_persistence"
)

func main() {
	zips := []string{"94103", "94110"}

	for _, zip := range(zips) {
		conditions, err := fetcher.FetchConditions(zip)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			conditions_persistence.StoreConditions(conditions)
		}
	}
}