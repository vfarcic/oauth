package main

import (
	"testing"
	"labix.org/v2/mgo/bson"
	"github.com/stretchr/testify/assert"
)

var testUser = getMongoUser(GetTestUser())

func TestSaveToMongoDBShouldInsertData(t *testing.T) {
	DropFromDB()
	SaveToMongoDB(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := MongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).One(&actual)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestSaveToMongoDBShouldUpdateWhenEmailExists(t *testing.T) {
	DropFromDB()
	SaveToMongoDB(testUser)
	SaveToMongoDB(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []MongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).All(&actual)

	assert.Nil(t, err)
	assert.Len(t, actual, 1)
}

func TestGetFromDBShouldReturnData(t *testing.T) {
	DropFromDB()
	SaveToMongoDB(testUser)

	actual, err := GetFromDB(testUser.Email)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestGetFromDBShouldReturnErrorWhenNonExistent(t *testing.T) {
	_, err := GetFromDB("john.doe@gmail.com")

	assert.NotNil(t, err)
}

func TestGetFromDBByAuthIDShouldReturnData(t *testing.T) {
	DropFromDB()
	SaveToMongoDB(testUser)

	actual, err := GetFromDBByAuthID(testUser.AuthID)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestGetFromDBByAuthIDShouldReturnErrorWhenNonExistent(t *testing.T) {
	_, err := GetFromDBByAuthID("111111111111111")

	assert.NotNil(t, err)
}

func TestDropFromDBShouldRemoveCollection(t *testing.T) {
	DropFromDB()
	SaveToMongoDB(testUser)

	err := DropFromDB()

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []MongoUser{}
	c.Find(nil).All(&actual)
	assert.Nil(t, err)
	assert.Len(t, actual, 0)
}
