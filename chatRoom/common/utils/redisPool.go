package utils

//redis连接池

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	redisURL         = "redis://0.0.0.0:6379"
	redisMaxIdle     = 50  //最大空闲连接数
	redisMaxActive   = 0   //最大的连接数，0不限制
	redisIdleTimeout = 100 //最大空闲连接时间
	redisPassword    = ""
)

var RedisPoolDb *redis.Pool

// redisPool 返回redis连接池
func RedisPool() *redis.Pool {
	RedisPoolDb = &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeout * time.Second,
		MaxActive:   redisMaxActive, //最大的连接数，0不限制
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if redisPassword != "" {
				if _, authErr := conn.Do("AUTH", redisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
	return RedisPoolDb
}
