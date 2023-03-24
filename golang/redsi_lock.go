package tools


import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisAddr = "localhost:6379"

type RedisLock struct {
	Pool     *redis.Pool
	LockName string
	Token    string
}

func (l *RedisLock) GetLock() bool {
	conn := l.Pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("SET", l.LockName, "NX", "EX", time.Second*10))
	if err == nil && result == "OK" {
		return true
	}
	return false
}

func (l *RedisLock) UnLock() {
	conn := l.Pool.Get()
	defer conn.Close()
	conn.Do("DEL", l.LockName)
}

func NewRedsiLock(pool *redis.Pool, name string) *RedisLock {
	return &RedisLock{
		Pool:     pool,
		LockName: name,
		Token:    time.Now().String(),
	}
}

func Example() {
	pool := &redis.Pool{
		MaxIdle:         5,
		MaxActive:       10,
		IdleTimeout:     10 * 60 * time.Second,
		MaxConnLifetime: 10,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisAddr)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	lock := NewRedsiLock(pool, "lock-test")
	if lock.GetLock() {

		// 释放锁
		lock.UnLock()
	} else {

	}
}
