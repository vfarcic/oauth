package main

import "github.com/stretchr/testify/mock"

type TestDB struct {
	mock.Mock
}

func (m TestDB) Save(user MongoUser) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m TestDB) Get(email string) (MongoUser, error) {
	ret := m.Called(email)
	return ret.Get(0).(MongoUser), ret.Error(1)
}

func (m TestDB) GetByAuthID(authID string) (MongoUser, error) {
	ret := m.Called(authID)
	return ret.Get(0).(MongoUser), ret.Error(1)
}
