package configuration

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "TASK_API"
)

// Config is the envconfig-compatible configuration struct for this server. See https://github.com/kelseyhightower/envconfig for more detail
type Config struct {
	Port      int    `envconfig:"PORT" default:"8080"`
	MongoHost string `envconfig:"MONGO_HOST" default:"localhost"`
}

// GetConfig uses envconfig to populate and return a Config struct. Returns all envconfig errors if they occurred
func GetConfig() (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(AppName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
