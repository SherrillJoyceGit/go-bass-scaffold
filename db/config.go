package db

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
