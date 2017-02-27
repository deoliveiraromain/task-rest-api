package configuration

import (
	"io/ioutil"
	"encoding/json"
)


// Config is the envconfig-compatible configuration struct for this server. See https://github.com/kelseyhightower/envconfig for more detail
type Config struct {
	Port      string `json:"ApiPort"`
	MongoHost string `json:"MongoHost"`
}

// GetConfig uses envconfig to populate and return a Config struct. Returns all envconfig errors if they occurred
func GetConfig() (*Config, error) {

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	if err = json.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
