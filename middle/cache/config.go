package cache

import "errors"

type Config struct {
	Addr     string // "localhost:6379"
	Password string
	Port     string
	Db       int
}

var CurrentConfig = Config{}

func getCurrentConfig() (*Config, error) {
	if CurrentConfig.Addr == "" ||
		(CurrentConfig.Db >= 0 && CurrentConfig.Db < 16) ||
		CurrentConfig.Password == "" ||
		CurrentConfig.Port == "" {

		return nil, errors.New("current config is not be settled ! ")

	} else {
		return &CurrentConfig, nil
	}
}
