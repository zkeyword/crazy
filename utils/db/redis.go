package db

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

// StartRedis 初始化redis
func StartRedis(addr string, password string, db, maxIdle, maxOpen int) (err error) {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxOpen,
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(db), redis.DialPassword(password))
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}

	conn := GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed")
	}

	return
}

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// GetRedisPool
func GetRedisPool() *redis.Pool {
	return redisPool
}

// 关闭redis
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}

// SetKey 设置key
func SetKey(key string, value string) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// GetKey 获取key
func GetKey(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return value, nil
}
