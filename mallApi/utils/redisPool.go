package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golangPractise/mallApi/config"
	"time"
)

// RedisPool 返回redis连接池
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.RedisMaxIdle,
		IdleTimeout: config.RedisIdleTimeoutSec * time.Second,
		MaxActive:   config.RedisMaxActive, //最大的连接数，0不限制
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(config.RedisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if config.RedisPassword != "" {
				if _, authErr := c.Do("AUTH", config.RedisPassword); authErr != nil {
					c.Close()
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}

			if config.RedisDbNum != 0 {
				_, err := c.Do("SELECT", config.RedisDbNum)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				c.Close()
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}
