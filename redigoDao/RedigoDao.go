package redigoDao

import (
	"github.com/garyburd/redigo/redis"

	"github.com/SwathiAR/redigoDao/connection"

	"fmt"
)

var redisConnPool *redis.Pool = connection.GetConnectionpool(connection.GetRedisConfig())

func Get(key interface{}) (interface{} , error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key));


}

func Set(key interface{}, vars interface{}) (bool , error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	isSet , err := redis.String(conn.Do("SET", key, vars))
	if(isSet== "OK"){
		return true , err
	}
	return  false , err

}

func Delete(key interface{}) (bool , error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	deleted, err := redis.Int(conn.Do("DEL", key))
         fmt.Println(deleted)
	if(deleted == 0){
		return false , err
	}
	return true , err
}

func SetIfNotExists(key interface{}, vars interface{})  (bool , error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	isSet, err := redis.Int(conn.Do("SETNX", key, vars))

	if (isSet == 0){
		return false , err
	}
	return  true , err
}

func ExpireKey(key interface{}, time int) (bool, error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	expires, err:= redis.Int(conn.Do("EXPIRE", key, time))


	if expires == 1 {
		return true , err
	}
	return false , err

}

func GetExpiryTime(key interface{}) (interface{} , error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	expires, err := redis.Int(conn.Do("TTL", key))


	if expires == -2 {
		return "The key doesn't exist anymore" , err
	}

	if expires == -1 {
		return "The key never expires" , err
	}

	return expires , err
}

func Persist(key interface{}) (bool , error){
	conn := redisConnPool.Get()
	defer conn.Close()
	persisted, err := redis.Int(conn.Do("PERSIST", key))
	fmt.Println(persisted)
	fmt.Println(err)

	if persisted == 0 {
		return  false , err
	}
	return true , err
}

func ContainsKey(key interface{}) (bool , error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	contains, err := redis.Int(conn.Do("EXISTS", key))


	if contains == 1 {
		return true , err
	}

	return false , err
}

func DeleteAllKeys () (interface{} , error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	return conn.Do("FLUSHALL")

}

