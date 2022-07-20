package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*type Config struct {
	DbType, DbName, User, Password, Host, Port string
}

type RedisConfig struct {
	Addr     string // "localhost:6379"
	Password string
	Port     string
	Db       int
}

type LogConfig struct {
	Net       string
	Host      string
	ServiceId string
}

type Swagger struct {
	Init int
}

var DbConfig Config

var AuthRedisConfig RedisConfig

var LogStashConfig LogConfig

var SwaggerConfig Swagger*/

func NewConfig() {

	// 预设命令行参数
	// 运行模式，用来指定运行的配置文件
	pflag.String("mode", "dev", "set config mode")
	// 运行端口
	pflag.Int("port", 8090, "set the port app run on")
	// gin运行模式
	pflag.String("ginMode", "release", "set the gin mode to run")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	mode := viper.GetString("mode")
	switch mode {
	case "dev":
		viper.SetConfigName("config_dev")
	case "release":
		viper.SetConfigName("config_release")
	case "test":
		viper.SetConfigName("config_test")
	default:
		// 默认为开发环境的配置
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			panic(fmt.Errorf("Unexpected error : %s \n", err))
		}
	}

	// 设置 Cloud 配置
	/*	if convertErr := viper.Sub("db").Sub("cloud").Unmarshal(&DbConfig); convertErr != nil {
			panic(fmt.Errorf("Convert config value error : %s \n", convertErr))
		} else {
			log.Printf("connect to %s for Cloud", DbConfig.Host)
			log.Printf("Convert Cloud config Ok !")
		}*/

	// 设置 redis 配置
	/*	if convertErr := viper.Sub("redis").Unmarshal(&AuthRedisConfig); convertErr != nil {
			panic(fmt.Errorf("Convert redis config value error : %s \n", convertErr))
		} else {
			log.Printf("Convert redis config Ok !")
		}*/

	// 设置 log 配置
	/*	if convertErr := viper.Sub("log").Unmarshal(&LogStashConfig); convertErr != nil {
			panic(fmt.Errorf("Convert log config value error : %s \n", convertErr))
		} else {
			log.Printf("Convert log config Ok !")
		}*/

	//swagger配置
	/*	if convertErr := viper.Sub("swagger").Unmarshal(&SwaggerConfig); convertErr != nil {
			panic(fmt.Errorf("Convert swagger config value error : %s \n", convertErr))
		} else {
			log.Printf("Convert swagger config Ok !")
		}*/
}
