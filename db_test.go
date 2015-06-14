package main

import (
	"testing"
	"labix.org/v2/mgo/bson"
	"github.com/stretchr/testify/assert"
)

var db = MongoDB{}
var testUser = GetMongoUser(GetTestUser())

// Save

func TestSaveShouldInsertData(t *testing.T) {
	db.Drop()
	db.Save(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := MongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).One(&actual)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestSaveShouldUpdateWhenEmailExists(t *testing.T) {
	db.Drop()
	db.Save(testUser)
	db.Save(testUser)

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []MongoUser{}
	err := c.Find(bson.M{"email": testUser.Email}).All(&actual)

	assert.Nil(t, err)
	assert.Len(t, actual, 1)
}

func TestGetShouldReturnData(t *testing.T) {
	db.Drop()
	db.Save(testUser)

	actual, err := db.Get(testUser.Email)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

// Get

func TestGetShouldReturnErrorWhenNonExistent(t *testing.T) {
	_, err := db.Get("john.doe@gmail.com")

	assert.NotNil(t, err)
}

// GetByAuthID

func TestGetByAuthIDShouldReturnData(t *testing.T) {
	db.Drop()
	db.Save(testUser)

	actual, err := db.GetByAuthID(testUser.AuthID)

	assert.Nil(t, err)
	assert.Equal(t, testUser, actual)
}

func TestGetByAuthIDShouldReturnErrorWhenNonExistent(t *testing.T) {
	_, err := db.GetByAuthID("111111111111111")

	assert.NotNil(t, err)
}

// Drop

func TestDropShouldRemoveCollection(t *testing.T) {
	db.Drop()
	db.Save(testUser)

	err := db.Drop()

	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	actual := []MongoUser{}
	c.Find(nil).All(&actual)
	assert.Nil(t, err)
	assert.Len(t, actual, 0)
}
