package redigoDao

import (
	"testing"
	"time"
)



func TestConfig(t *testing.T){

expectedPassword:= "Cisco123!"
expectedPort:= "6379"
expectedAddress:= "localhost"
expectedProtocol:= "tcp"
	redisConf:= GetRedisConfig()
	if(redisConf.address != expectedAddress){
		t.Error("Actual: " +  redisConf.address + "Expected: " + expectedAddress )
	}

	if(redisConf.port != expectedPort){
		t.Error("getPort failed")
	}

	if(redisConf.password != expectedPassword){
		t.Error("getPassword failed")
	}

	if(redisConf.protocol != expectedProtocol){
		t.Error("getProtocol failed")
	}


}

func TestGetConnectionpool(t *testing.T) {
	expectedIdleTimeout:= 300 * time.Second
	expectedMaxIdle := 3

	redisConnPool := GetConnectionpool(GetRedisConfig())
	conn1:= redisConnPool.Get()
	conn2:= redisConnPool.Get()
	conn3:= redisConnPool.Get()
	conn4:= redisConnPool.Get()

	defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()
	defer conn4.Close()


	if(conn1.Err()!= nil){
		t.Error("Unable to get 1st connection")

	}

	if(conn2.Err()!= nil){
		t.Error("Unable to get 2nt connection")

	}

	if(conn3.Err()!= nil){
		t.Error("Unable to get 3rt connection")

	}

	if(conn4.Err() == nil){
		t.Error("MaxActive is not working")
	}



	if(redisConnPool.MaxIdle != expectedMaxIdle){
		t.Error("MaxIdle is not set properly")

	}
	if(redisConnPool.IdleTimeout != expectedIdleTimeout){
		t.Error("IdleTimeout is not set properly")

	}
}


