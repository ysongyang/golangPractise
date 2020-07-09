package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func RedisInit() redis.Conn {
	network := "tcp"
	address := "localhost:6379"
	RedisPassword := ""
	redisPool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //最大的连接数，0不限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接池代码
			c, err := redis.Dial(network, address)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			if RedisPassword != "" {
				//验证redis密码
				if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			return c, nil
		},
	}
	return redisPool.Get()
}
