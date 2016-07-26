package redisMap

import (
	"github.com/garyburd/redigo/redis"

	"github.com/SwathiAR/redigoDao/connection"
)

var redisConnPool *redis.Pool = connection.GetConnectionpool(connection.GetRedisConfig())

func SetMapField(key, field, value string) (bool, error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("HSET", key, field, value))

}

func GetMapField(key, field string) (string, error) {
	conn := redisConnPool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("HGET", key, field))

	return value, err

}

func DeleteMapField(key, field string) (bool, error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("HDEL", key, field))
}

func GetMap(key string) (interface{}, error) {
	conn := redisConnPool.Get()
	defer conn.Close()

	value, err := redis.StringMap(conn.Do("HGETALL", key))

	return value, err
}

func GetFields (key string) (interface{} , error) {

	conn := redisConnPool.Get()
	defer conn.Close()


	return redis.Strings(conn.Do("HKEYS" , key))
}
































//func SetMap (key string ,mymap map[string]string) (bool , error){
//
//	conn:= redisConnPool.Get()
//
//	isSet , err:=redis.String(conn.Do("HMSET" , key , MapToSlice(mymap)))
//
//	fmt.Println(isSet)
//	fmt.Println(err)
//	if(isSet=="OK"){
//		return true , err
//	}
//
//	return  false , err
//
//
//
//
//
//}
//
//
//func MapToSlice(m map[string]string) ([]string){
//
//	var resultantSlice = make([]string , len(m)*2)
//
//	i:=0
//	for k , v := range m {
//		resultantSlice[i] = k
//		resultantSlice[i+1] = v
//		i = i+2
//	}
//	fmt.Println(resultantSlice)
//	return resultantSlice
//
//}