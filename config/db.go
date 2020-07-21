package config

// DbConf Mysql 数据库结构
type DbConf struct {
	Dsn     string
	MaxIdle int
	MaxOpen int
}

// DbConfig  Mysql主库配置
var DbConfig DbConf = DbConf{
	Dsn:     "root:123456@tcp(127.0.0.1:3306)/b?charset=utf8mb4&parseTime=true",
	MaxIdle: 10,
	MaxOpen: 100,
}

// RedisDbConf Redis数据库结构
type RedisDbConf struct {
	Addr     string
	Password string
	DB       int
	MaxIdle  int
	MaxOpen  int
}

// RedisDbConfig Redis 数据库结构
var RedisDbConfig RedisDbConf = RedisDbConf{
	Addr:     "127.0.0.1:6379",
	DB:       1,
	Password: "123456",
	MaxIdle:  10,
	MaxOpen:  100,
}
