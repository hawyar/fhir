package server

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	client, err := redis.Dial("tcp", "redis:6379")

	if err != nil {
		log.Fatal(err)
	}

	return &redis.Pool{
		MaxIdle:   3,
		MaxActive: 10,
		Dial: func() (redis.Conn, error) {
			return client, nil
		},
	}
}

func Set(key string, val string) {
	conn := pool.Get()
	ok, err := conn.Do("SET", key, val)

	if err != nil {
		panic(err.Error())
	}

	if ok != "OK" {
		panic("SET failed")
	}

	defer conn.Close()
}

func Get(key string) string {
	conn := pool.Get()
	val, err := redis.String(conn.Do("GET", key))

	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	return val
}
