package redigoDao

import (
	"testing"
	"github.com/stretchr/testify/assert"

)



func TestSet(t *testing.T) {
	isSet , err:=Set("serialnumber", 1000)
	assert.True(t , isSet)
	assert.Nil(t , err)

	reply , _:= Get("serialnumber")
	assert.NotNil(t, reply)
	assert.Equal(t, reply, "1000")

	defer DeleteAllKeys()


}

func TestGet(t *testing.T) {
	Set("serialnumber" , 1000)

	result , err := Get("serialnumber")
	assert.Equal(t, "1000", result)
	assert.Nil(t , err )

	result1 , err1 := Get("phone")
	assert.Equal(t, result1, "")
	assert.NotNil(t , err1)
	defer DeleteAllKeys()
}

func TestDelete(t *testing.T) {
	deleted, err:=Delete("name")
	assert.False(t , deleted)
	assert.Nil(t , err)

	result , _:= Get("name")
	assert.Equal(t, result, "")


	Set("serialnumber" , 1000)

	deleted1, err1:=Delete("serialnumber")
	assert.True(t , deleted1)
	assert.Nil(t , err1)

	result1 , _:= Get("serialnumber")
	assert.Equal(t, result1, "")

        defer DeleteAllKeys()

}

func TestSetIfNotExists(t *testing.T) {

	isSet , err:= SetIfNotExists("name", "swathi")
	assert.True(t , isSet)
	assert.Nil(t ,err)

	result , _:= Get("name")
	assert.Equal(t, "swathi", result)

	isSet1 , err1:= SetIfNotExists("name", "swrathna")
	assert.False(t , isSet1)
	assert.Nil(t ,err1)
	result1 ,_:= Get("name")
	assert.Equal(t, "swathi", result1)

	defer DeleteAllKeys()

}

func TestExpireAKey(t *testing.T) {
	Set("serialnumber" , 1000)
	expires,err:=ExpireKey("serialnumber", 50)
	assert.True(t , expires)
	assert.Nil(t ,err)

	et1,_:= GetExpiryTime("serialnumber")
	assert.Equal(t, et1, 50)
	assert.NotEqual(t, et1, "The key never expires")
	assert.NotEqual(t, et1, "The key doesn't exist anymore")

	expires1,err1:=ExpireKey("name", 50)
	assert.False(t , expires1)
	assert.Nil(t ,err1)

	defer DeleteAllKeys()


}

func TestGetExpiryTime(t *testing.T) {
	Set("serialnumber" , 1000)
	eTime,err:= GetExpiryTime("serialnumber")
	assert.Equal(t, eTime, "The key never expires")
	assert.Nil(t ,err)

        ExpireKey("serialnumber", 400)
	eTime2, err2 := GetExpiryTime("serialnumber")
	assert.Equal(t, eTime2, 400)
	assert.NotEqual(t, eTime2, "The key never expires")
	assert.NotEqual(t, eTime2, "The key doesn't exist anymore")
	assert.Nil(t ,err2)

	defer DeleteAllKeys()

}

func TestPersist(t *testing.T) {
	Set("connections" , 10)
	isPersistent,err:=Persist("connections")
	assert.False(t , isPersistent)
	assert.Nil(t ,err)
	eTime,_:= GetExpiryTime("connections")
	assert.Equal(t , eTime , "The key never expires")

	ExpireKey("connections" , 80)
	isPersistent1 , err1:=Persist("connections")
	assert.True(t , isPersistent1)
	assert.Nil(t ,err1)

	eTime2,_:= GetExpiryTime("connections")
	assert.Equal(t , eTime2 , "The key never expires")

	defer DeleteAllKeys()
}

func TestContainsKey(t *testing.T) {

	Set("connections" , 10)

	hasKey,err:= ContainsKey("connections")
	assert.True(t , hasKey)
	assert.Nil(t ,err)

	hasKey1,err1:= ContainsKey("profile")
	assert.False(t , hasKey1)
	assert.Nil(t ,err1)

	defer DeleteAllKeys()
}


