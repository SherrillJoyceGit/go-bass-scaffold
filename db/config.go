package db

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Config defines the config of db something
type Config struct {
	DbType   string
	DbName   string
	UserName string
	Password string
	Host     string
	Port     string
}

/*var CurrentConfig = Config{}

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
}*/

func NewFileConfig(vp *viper.Viper) (*Config, error) {
	cfg := new(Config)
	if convertErr := vp.Sub("db").Unmarshal(&cfg); convertErr != nil {
		panic(fmt.Errorf("Convert database config from config file  error : %s \n", convertErr))
	} else {
		log.Printf("Convert database config from config file successfully !")
		return cfg, nil
	}

}
