package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func init() {
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
}

func main() {
	conn := redisPool.Get()
	defer conn.Close()
	/*_, err = c.Do("SET", "mykey", "superWang往")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, _ = c.Do("HSET", "user01", "name", "张三")
	_, _ = c.Do("HSET", "user01", "age", "30")

	hget, err := redis.String(c.Do("HGET", "user01", "name"))
	if err != nil {
		fmt.Println("redis HGET failed:", err)
	} else {
		fmt.Printf("HGET user01 name: %v \n", hget)
	}

	age, err := redis.String(c.Do("HGET", "user01", "age"))
	if err != nil {
		fmt.Println("redis HGET failed:", err)
	} else {
		fmt.Printf("HGET user01 age: %v \n", age)
	}

	_, _ = c.Do("HMSET", "user02", "name","zhangsan","age","33")

	str, err := redis.Strings(c.Do("HGETALL", "user02"))
	if err != nil {
		fmt.Println("redis HMGET failed:", err)
	} else {
		fmt.Println(str)
	}*/

	_, _ = conn.Do("HMSET", "user02", "name", "zhangsan", "age", "33")

	//HMGET 返回切片
	str, err := redis.Strings(conn.Do("HMGET", "user02", "name", "age"))
	if err != nil {
		fmt.Println("redis HMGET failed:", err)
	} else {
		fmt.Println(str)
	}
	for _, val := range str {
		fmt.Println(val)
	}
	/*str, err := redis.Strings(c.Do("HGETALL", "user02"))
	if err != nil {
		fmt.Println("redis HMGET failed:", err)
	} else {
		fmt.Println(str)
	}*/
}
