package db

import "errors"

// Config defines the config of db something
type Config struct {
	DbType   string
	DbName   string
	Username string
	Password string
	Host     string
	Port     string
}

var CurrentConfig = Config{}

func getCurrentConfig() (*Config, error) {
	if CurrentConfig.DbType == "" ||
		CurrentConfig.DbName == "" ||
		CurrentConfig.Username == "" ||
		CurrentConfig.Password == "" ||
		CurrentConfig.Host == "" ||
		CurrentConfig.Port == "" {

		return nil, errors.New("current config is not be settled ! ")

	} else {
		return &CurrentConfig, nil
	}
}
