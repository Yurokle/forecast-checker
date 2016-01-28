package fetcher

import (
	"fmt"
	"io/ioutil"
//	"encoding/json"
	"net/http"
)


const key = "c33b9575f02fa950"

func conditionsURL() string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/q/94103.json", key)
}

func FetchConditions() (string, error) {
	resp, err := http.Get(conditionsURL())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	jsonResponse := string(rawBody)

	return jsonResponse, nil
}
