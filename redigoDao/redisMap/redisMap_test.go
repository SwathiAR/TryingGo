package redisMap

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/SwathiAR/redigoDao"
)

func TestSetMap(t *testing.T) {
	isSet , err := SetMapField("user1" , "firstName" , "xyz")
	assert.True(t ,isSet)
	assert.Nil(t ,err)

	value, err:= GetMapField("user1" , "firstName")
	assert.Equal(t , value , "xyz")
	assert.Nil(t , err )

	defer redigoDao.DeleteAllKeys()

}

func TestGetMapField(t *testing.T) {
	//var user1 = map[string]string {
	//	"firstName" : "xyz" ,
	//	"secondName" : "pqr" ,
	//	"email" : "email@example.com",
	//
	//}

	SetMapField("user1" , "firstName" , "xyz")
	SetMapField("user1" , "secondName" , "pqr")
	SetMapField("user1" , "email" , "email@example.com")


	value, err:= GetMapField("user1" , "secondName")
	assert.Equal(t , value , "pqr")
	assert.Nil(t , err )


	value1, err1:= GetMapField("user" , "name")
	assert.Equal(t , value1 , "")
	assert.NotNil(t , err1 )

	defer redigoDao.DeleteAllKeys()

}


func TestDeleteMapField(t *testing.T) {

	SetMapField("user1" , "firstName" , "xyz")

	deleted , err := DeleteMapField("user1" , "firstName")
	assert.True(t , deleted)
	assert.Nil(t , err)

	value, err1:= GetMapField("user1" , "firstName")
	assert.Equal(t , value , "")
	assert.NotNil(t ,err1)

	defer redigoDao.DeleteAllKeys()
}


func TestGetMap(t *testing.T) {

	SetMapField("user1" , "firstName" , "xyz")
	SetMapField("user1" , "secondName" , "pqr")
	SetMapField("user1" , "email" , "email@example.com")


	value, err := GetMap("user1")
	assert.Equal(t , map[string]string{"firstName":"xyz", "secondName":"pqr", "email":"email@example.com"} ,value )
	assert.Nil(t ,err)

	defer redigoDao.DeleteAllKeys()

}


func TestGetFields(t *testing.T) {
	SetMapField("user1" , "firstName" , "xyz")
	SetMapField("user1" , "secondName" , "pqr")
	SetMapField("user1" , "email" , "email@example.com")

	fields,err:= GetFields("user1")
	assert.Equal(t , fields , []string{"firstName" , "secondName" , "email"})
	assert.Nil(t , err)

	defer redigoDao.DeleteAllKeys()

}