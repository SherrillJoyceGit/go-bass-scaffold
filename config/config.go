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

// NewConfig 根据yaml配置文件不同段落的存在情况，进行配置注入
//func NewConfig() {

// 预设命令行参数
// 运行模式，用来指定运行的配置文件
/*	pflag.String("configName", "config", "set config file name, only for .yaml file now")
	// 运行端口
	//pflag.Int("port", 8090, "set the port app run on")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	fn := viper.GetString("configName")
	viper.SetConfigName(fn)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			panic(fmt.Errorf("Unexpected error : %s \n", err))
		}
	}

	// 设置 db 配置
	if convertErr := viper.Sub("db").Unmarshal(&db.CurrentConfig); convertErr != nil {
		panic(fmt.Errorf("Convert config value error : %s \n", convertErr))
	} else {
		log.Printf("connect to %s for db", db.CurrentConfig.Host)
		log.Printf("Convert db config Ok !")
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
//}

func NewViper() *viper.Viper {
	// 预设命令行参数
	// 运行模式，用来指定运行的配置文件
	pflag.String("configName", "config", "set config file name, only for .yaml file now")
	// 运行端口
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	fn := viper.GetString("configName")
	viper.SetConfigName(fn)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		} else {
			panic(fmt.Errorf("Unexpected error : %s \n", err))
		}
	}

	return viper.GetViper()
}
