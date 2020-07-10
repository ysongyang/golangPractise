package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	RedisURL            = "redis://0.0.0.0:6379"
	RedisMaxIdle        = 8   //最大空闲连接数
	RedisMaxActive      = 0   //最大的连接数，0不限制
	RedisIdleTimeoutSec = 240 //最大空闲连接时间
	RedisPassword       = ""
)

// RedisPool 返回redis连接池
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     RedisMaxIdle,
		IdleTimeout: RedisIdleTimeoutSec * time.Second,
		MaxActive:   RedisMaxActive, //最大的连接数，0不限制
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(RedisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if RedisPassword != "" {
				if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
