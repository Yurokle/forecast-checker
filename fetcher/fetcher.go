package fetcher

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"forecast-checker/conditions"
)


const key = "c33b9575f02fa950"

func conditionsURL(zip string) string {
	return fmt.Sprintf("http://api.wunderground.com/api/%s/conditions/q/%s.json", key, zip)
}

func fetchConditionsJson(zip string) (*[]byte, error) {
	resp, err := http.Get(conditionsURL(zip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw_json, err := ioutil.ReadAll(resp.Body)
	return &raw_json, err
}

func FetchConditions(zip string) (*conditions.Conditions, error) {
	raw_json, err := fetchConditionsJson(zip)
	if err != nil {
		return nil, err
	}

	var data conditions.Conditions
	err = json.Unmarshal(*raw_json, &data)
	data.Json = string(*raw_json)

	return &data, err
}
