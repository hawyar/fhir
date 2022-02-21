package main

import (
	"github.com/gomodule/redigo/redis"
)

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "fhir_redis_1:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
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
