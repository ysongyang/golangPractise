package config

const (
	//mysql数据库配置

	DriverName = "mysql"
	USERNAME   = "root"
	PASSWORD   = "root"
	NETWORK    = "tcp"
	SERVER     = "localhost"
	PORT       = 3306
	DATABASE   = "golang"

	//redis数据库配置

	RedisURL            = "redis://0.0.0.0:6379"
	RedisMaxIdle        = 256 //最大空闲连接数
	RedisMaxActive      = 0   //最大的连接数，0不限制
	RedisIdleTimeoutSec = 240 //最大空闲连接时间
	RedisPassword       = ""
	RedisDbNum          = 2 //redis数据库
)
