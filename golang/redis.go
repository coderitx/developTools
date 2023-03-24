package tools


import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisClient represents a Redis client
type RedisClient struct {
	pool *redis.Pool
}

// NewRedisClient creates a new RedisClient
func NewRedisClient(server, password string, db int) *RedisClient {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if db != 0 {
				if _, err := c.Do("SELECT", db); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
	return &RedisClient{pool}
}

// Close closes the RedisClient
func (c *RedisClient) Close() error {
	return c.pool.Close()
}

// GetString gets a string value from Redis at the specified key
func (c *RedisClient) GetString(key string) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// SetString sets a string value in Redis at the specified key
func (c *RedisClient) SetString(key string, value string) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	return err
}

// GetInt gets an integer value from Redis at the specified key
func (c *RedisClient) GetInt(key string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

// SetInt sets an integer value in Redis at the specified key
func (c *RedisClient) SetInt(key string, value int) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	return err
}

// Exists checks if a key exists in Redis
func (c *RedisClient) Exists(key string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

// Delete deletes a key in Redis
func (c *RedisClient) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

// Incr increments a value in Redis at the specified key
func (c *RedisClient) Incr(key string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

// Decr decrements a value in Redis at the specified key
func (c *RedisClient) Decr(key string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DECR", key))
}
