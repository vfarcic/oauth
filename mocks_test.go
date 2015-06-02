package main

import (
	"github.com/stretchr/gomniauth/test"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

func GetTestUser() *test.TestUser {
	testUser := new(test.TestUser)
	testUser.On("Email").Return("viktor@farcic.com")
	testUser.On("Name").Return("Viktor")
	testUser.On("Nickname").Return("vfarcic")
	testUser.On("AvatarURL").Return("http://mygravatar.com")
	testUser.On("AuthCode").Return("123")
	testUser.On("ProviderCredentials").Return(make(map[string]*common.Credentials))
	testUser.On("Data").Return(make(objx.Map))
	return testUser
}
