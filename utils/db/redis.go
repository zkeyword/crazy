// package db

package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

// StartRedis 初始化 Redis 连接池
func StartRedis(addr, password string, db, maxIdle, maxOpen int) error {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxOpen,
		IdleTimeout: 30 * time.Minute,
		Dial: func() (redis.Conn, error) {
			// Dial 不会返回非 nil conn + error，所以无需 c.Close()
			return redis.Dial(
				"tcp",
				addr,
				redis.DialDatabase(db),
				redis.DialPassword(password),
			)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	// 测试连接
	conn := GetRedis()
	if conn == nil {
		return errors.New("failed to get redis connection")
	}
	defer conn.Close()

	pong, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}
	if pong != "PONG" {
		return errors.New("redis ping returned non-PONG response")
	}

	return nil
}

// GetRedis 获取连接，安全检查
func GetRedis() redis.Conn {
	if redisPool == nil {
		return nil
	}
	return redisPool.Get()
}

// GetRedisPool 获取连接池（供高级使用）
func GetRedisPool() *redis.Pool {
	return redisPool
}

// CloseRedis 关闭连接池
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}

// SetKey 设置 key（永不过期）
func SetKey(key, value string) error {
	conn := GetRedis()
	if conn == nil {
		return errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}

// SetKeyEx 设置 key 并指定过期时间（秒）
func SetKeyEx(key, value string, expireSeconds int) error {
	conn := GetRedis()
	if conn == nil {
		return errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", expireSeconds)
	return err
}

// GetKey 获取 key
func GetKey(key string) (string, error) {
	conn := GetRedis()
	if conn == nil {
		return "", errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", nil // key 不存在，返回空字符串 + nil error（按需可改为自定义错误）
	}
	return value, err
}

// Exists 检查 key 是否存在
func Exists(key string) (bool, error) {
	conn := GetRedis()
	if conn == nil {
		return false, errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	return exists, err
}

// Del 删除 key
func Del(key string) error {
	conn := GetRedis()
	if conn == nil {
		return errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

// TTL 获取 key 剩余生存时间（秒），-1=永不过期，-2=不存在
func TTL(key string) (int64, error) {
	conn := GetRedis()
	if conn == nil {
		return -2, errors.New("redis pool is not initialized")
	}
	defer conn.Close()

	ttl, err := redis.Int64(conn.Do("TTL", key))
	return ttl, err
}
