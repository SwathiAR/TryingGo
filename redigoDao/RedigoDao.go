package redigoDao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
	"log"
)

type redisConfig  struct {
	password string
	protocol string
	address  string
	port     string
}

func GetRedisConfig() (* redisConfig) {
	viper.SetConfigName("daoCloud")
	viper.AddConfigPath("/Users/swrathna/PlayGround/TryingGo/src/github.com/SwathiAR/RedigoDao/")
	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(err1)
	}
	rc:= new (redisConfig)

		rc.password =viper.GetString("redisPassword")
		rc.protocol =viper.GetString("redisNetworkProtocol")
		rc.address=viper.GetString("redisAddress")
		rc.port= viper.GetString("redisPort")
	return rc
}

func GetConnectionpool(rc *redisConfig) *redis.Pool {
	return &redis.Pool{
                MaxActive: 3,
		MaxIdle:3,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, errDial := redis.Dial(rc.protocol, ":" + rc.port)

			if errDial != nil {
				log.Fatalln("Redis dial failed", errDial)
				return nil, errDial
			}
			if _, errAuth := conn.Do("Auth", rc.password); errAuth != nil {
				log.Fatalln("Redis Authentication failed", errAuth)
				conn.Close()
				return nil, errAuth
			}
			log.Println("Authenticated")
			return conn, errDial
		},

	}

}

















