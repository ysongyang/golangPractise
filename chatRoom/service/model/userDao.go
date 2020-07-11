package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//服务器启动后，初始化userDao实例，做全局变量，需要redis操作时，直接使用即可

var (
	MyUserDao *UserDao
)

//定义UserDao 结构体
//完成对User结构体的各种操作，如：增删改查

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式
func NewUserDao(redisPool *redis.Pool) (userDao *UserDao) {
	return &UserDao{redisPool}
}

//根据用户id返回一个User实例和 error
func (userDao *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	//这里需要redis.string 进行转换
	res, err := redis.String(conn.Do("HGET", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXIST
		}
		return
	}
	user = &User{}
	//把res反序列化成User对象
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		//err = fmt.Errorf("getUserById json.Unmarshal error: %s", err)
		return
	}
	return user, nil
}

//完成登录的校验
//如果用户的id和pwd都是正确的 返回一个user实例
//如果用户的id或密码有错误，返回错误信息
func (userDao *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//根据id 到 redis 里查询
	conn := userDao.pool.Get()
	defer conn.Close()
	user, err = userDao.getUserById(conn, userId)

	if err != nil {
		return nil, err
	}

	//校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PASSWORD
		return
	}
	return user, nil
}

//注册的处理
func (userDao *UserDao) Register(user *User) (err error) {
	//根据id 到 redis 里查询
	conn := userDao.pool.Get()
	defer conn.Close()
	//校验id是否存在
	_, err = userDao.getUserById(conn, user.UserId)

	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//说明id不存在可以进行注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		errText := fmt.Sprintf("%s:%s", "register error", err)
		return errors.New(errText)
	}
	return
}
