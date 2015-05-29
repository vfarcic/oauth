package main

import (
	"testing"
	"labix.org/v2/mgo/bson"
	"github.com/stretchr/testify/assert"
)

var testUser = getMongoUser(GetTestUser())

func TestSaveToDBShouldInsertData(t *testing.T) {
	DropFromDB()
	SaveToDB(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := mongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).One(&actual)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestSaveToDBShouldUpdateWhenEmailExists(t *testing.T) {
	DropFromDB()
	SaveToDB(testUser)
	SaveToDB(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []mongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).All(&actual)

	assert.Nil(t, err)
	assert.Len(t, actual, 1)
}

func TestGetFromDBShouldReturnData(t *testing.T) {
	DropFromDB()
	SaveToDB(testUser)

	actual, err := GetFromDB(testUser.Email)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestGetFromDBShouldReturnErrorWhenNonExistent(t *testing.T) {
	_, err := GetFromDB("john.doe@gmail.com")

	assert.NotNil(t, err)
}

func TestDropFromDBShouldRemoveCollection(t *testing.T) {
	DropFromDB()
	SaveToDB(testUser)

	err := DropFromDB()

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []mongoUser{}
	c.Find(nil).All(&actual)
	assert.Nil(t, err)
	assert.Len(t, actual, 0)
}
