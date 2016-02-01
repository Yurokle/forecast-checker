package main

import (
	"fmt"
	"forecast-checker/fetcher"
)

func main() {
	conditions, err := fetcher.FetchConditions("94103")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("JSON:")
		fmt.Println(conditions.Current_observation)
		fmt.Println("RAW JSON:")
		fmt.Println(conditions.Json)
	}
}