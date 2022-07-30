package database

import (
	"github.com/gomodule/redigo/redis"
	"moell/config"
	"time"
)

var Connect *redis.Pool

func Open()  {
	Connect = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", config.Get().RedisHost)
			if err != nil {
				return nil, err
			}

			if config.Get().RedisPassword != "" {
				if _, err := c.Do("AUTH", config.Get().RedisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         config.Get().RedisMaxIdle,
		MaxActive:       config.Get().RedisMaxIdle,
		IdleTimeout:     config.Get().RedisIdleTimeout,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func Set(key string, value interface{}, time int) error {
	conn := Connect.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := Connect.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := Connect.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := Connect.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}


