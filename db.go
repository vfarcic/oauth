package main

import (
	"github.com/stretchr/gomniauth/common"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type DB interface {
	Save(user MongoUser) error
	Get(email string) (MongoUser, error)
	GetByAuthID(authID string) (MongoUser, error)
}

type MongoDB struct {
}

func (db MongoDB) Save(user MongoUser) error {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	_, err := c.Upsert(bson.M{"email": user.Email}, user)
	return err
}

func (db MongoDB) Get(email string) (MongoUser, error) {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	users := MongoUser{}
	err := c.Find(bson.M{"email": email}).One(&users)
	return users, err
}

func (db MongoDB) GetByAuthID(authID string) (MongoUser, error) {
	session := getSession()
	defer session.Close()
	c := getUsersCollection(session)
	users := MongoUser{}
	err := c.Find(bson.M{"authid": authID}).One(&users)
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
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session
}

type MongoUser struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Nickname string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	ProviderCredentials map[string]*common.Credentials `json:"provider_credentials"`
	AuthCode string `json:"auth_code"`
	AuthID string `json:"auth_id"`
}

func GetMongoUser(user common.User) MongoUser {
	return MongoUser {
		Email: user.Email(),
		Name: user.Name(),
		Nickname: user.Nickname(),
		AvatarURL: user.AvatarURL(),
		AuthCode: user.AuthCode(),
		ProviderCredentials: user.ProviderCredentials(),
		AuthID: user.Data().Get("id").Str(),
	}
}