package cache

type RedisConfig struct {
	Addr     string // "localhost:6379"
	Password string
	Port     string
	Db       int
}
