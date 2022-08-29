package gredis

import "github.com/gomodule/redigo/redis"

var RedisConn *redis.Pool

func Setup() {
	RedisConn = &redis.Pool{}
}
