package main

import (
	"forecast-checker/fetcher"
	"fmt"
)

func main() {
	response, err := fetcher.FetchConditions()

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println(response)
	}
}