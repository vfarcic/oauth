package main

import (
	"github.com/stretchr/gomniauth/common"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type mongoUser struct {
	Email string
	Name string
	Nickname string
	AvatarURL string
	ProviderCredentials map[string]*common.Credentials
	AuthCode string
}

func getMongoUser(user common.User) mongoUser {
	return mongoUser {
		Email: user.Email(),
		Name: user.Name(),
		Nickname: user.Nickname(),
		AvatarURL: user.AvatarURL(),
		AuthCode: user.AuthCode(),
		ProviderCredentials: user.ProviderCredentials(),
	}
}

func SaveToDB(user mongoUser) error {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	_, err := c.Upsert(bson.M{"email": user.Email}, user)
	return err
}

func GetFromDB(email string) (mongoUser, error) {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	users := mongoUser{}
	err := c.Find(bson.M{"email": email}).One(&users)
	return users, err
}

func DropFromDB() error {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	err := c.DropCollection()
	return err
}

func getUsersCollection(session *mgo.Session) *mgo.Collection {
	return session.DB("oauth").C("users")
}

func getSession() *mgo.Session {
	// Change to param
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session
}
