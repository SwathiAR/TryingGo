package redisSortedSet

import (
	"github.com/garyburd/redigo/redis"

	"github.com/SwathiAR/redigoDao/connection"
)

var redisConnPool *redis.Pool = connection.GetConnectionpool(connection.GetRedisConfig())

func AddSortedSet(key string, score int, value string) (bool, error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("ZADD", key, score, value))
}

func GetSortedSet(key string) (interface{}, error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("ZRANGEBYSCORE", key, "-inf", "+inf"))

}

func GetRangeOfSet(key string, min int, max int) (interface{}, error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	return redis.Strings(conn.Do("ZRANGE", key, min, max))

}