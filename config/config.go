package config

import(
	"os"
	"io/ioutil"
	"encoding/json"
)

type configMap map[string]string
var configJson configMap

func loadConfig(*configMap)  {
	file, err := os.Open("config.json")
	if err != nil {
		// Process Error
	}

	fileContent, err := ioutil.ReadAll(file)
	json.Unmarshal(fileContent, &configJson)
}

func GetConfigKey(key string) string{
	if configJson == nil {
		loadConfig(&configJson)
	}
	return configJson[key]
}
