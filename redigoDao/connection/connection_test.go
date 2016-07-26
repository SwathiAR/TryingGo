package connection

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)



func TestConfig(t *testing.T){

expectedPassword:= "Cisco123!"
expectedPort:= "6379"
expectedAddress:= "localhost"
expectedProtocol:= "tcp"
	redisConf:= GetRedisConfig()

	assert.Equal(t , expectedProtocol , redisConf.protocol)
	assert.Equal(t , expectedPassword , redisConf.password)
	assert.Equal(t , expectedAddress , redisConf.address)
	assert.Equal(t , expectedPort , redisConf.port)
}

func TestGetConnectionPool(t *testing.T) {
	expectedIdleTimeout:= 300 * time.Second
	expectedMaxIdle := 3
	expectedMaxActive := 3

	redisConnPool:= GetConnectionpool(GetRedisConfig())
	conn1:= redisConnPool.Get()
	conn2:= redisConnPool.Get()
	conn3:= redisConnPool.Get()
	conn4:= redisConnPool.Get()

	defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()
	defer conn4.Close()


	assert.Nil(t , conn1.Err())
	assert.Nil(t , conn2.Err())
	assert.Nil(t , conn3.Err())
	assert.NotNil(t , conn4.Err())

	assert.Equal(t ,redisConnPool.MaxIdle , expectedMaxIdle)
	assert.Equal(t , redisConnPool.IdleTimeout , expectedIdleTimeout)
        assert.Equal(t , redisConnPool.MaxActive , expectedMaxActive)


}


