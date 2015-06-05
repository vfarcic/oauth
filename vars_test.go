package main

import (
	"testing"
	"flag"
	"os"
	"github.com/stretchr/testify/assert"
	"log"
	"io/ioutil"
)

func TestGetFlagArgsShouldReturnTheSameName(t *testing.T) {
	beforeTest()
	expected := "name"

	name, _, _ := getFlagArgs(expected, "fullName", "defaultValue", false)
	if name != expected {
		t.Error("Return name should not be", name)
	}
}

func TestGetFlagArgsShouldReturnStringValue(t *testing.T) {
	beforeTest()

	_, value, _ := getFlagArgs("name", "fullName", "defaultValue", false)
	assert.NotEmpty(t, value)
}

func TestGetFlagArgsShouldReturnDefaultValue(t *testing.T) {
	beforeTest()
	expected := "default value"

	_, value, _ := getFlagArgs("name", "fullName", expected, false)
	assert.Equal(t, expected, value)
}

func TestGetFlagArgsShouldReturnEnvVariableInUpperCaseAsValue(t *testing.T) {
	beforeTest()
	name := "name"
	expected := "environment value"
	os.Setenv("NAME", expected)

	_, value, _ := getFlagArgs(name, "fullName", "defaultValue", false)
	assert.Equal(t, expected, value)
}

func TestGetFlagArgsShouldReplaceDashWithUnderscore(t *testing.T) {
	beforeTest()
	name := "my-variable"
	expected := "environment value"
	os.Setenv("MY_VARIABLE", expected)

	_, value, _ := getFlagArgs(name, "fullName", "defaultValue", false)
	assert.Equal(t, expected, value)
}

func TestGetFlagArgsShouldReturnUsage(t *testing.T) {
	beforeTest()
	name := "name"
	envName := "NAME"
	fullName := "This is the full name of the argument"
	expected := fullName + ". It can also be specified with environment variable " + envName + "."

	_, _, usage := getFlagArgs(name, fullName, "defaultValue", false)
	assert.Equal(t, expected, usage)
}

func TestGetFlagArgsShouldReturnUsageWithMandatory(t *testing.T) {
	beforeTest()
	name := "name"
	fullName := "This is the full name of the argument"
	expected := "This variable is mandatory."

	_, _, usage := getFlagArgs(name, fullName, "defaultValue", true)
	assert.Contains(t, usage, expected)
}

func TestGetProviderShouldContainAllData(t *testing.T) {
	clientId := "clientId"
	clientSecret := "clientSecret"
	redirectUrl := "redirectUrl"
	expected := &provider{
		clientId: clientId,
		clientSecret: clientSecret,
		redirectUrl: redirectUrl,
	}
	provider := getProvider("name", clientId, clientSecret, redirectUrl)
	assert.Equal(t, expected, provider)
}

func TestGetVarsShouldInvokeFlagUtilForDomain(t *testing.T) {
	expected := *mockedFlagUtil("domain", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.domain)
}

func TestGetVarsShouldInvokeFlagUtilForPort(t *testing.T) {
	expected := *mockedFlagUtil("port", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.port)
}

func TestGetVarsShouldInvokeFlagUtilForSecKey(t *testing.T) {
	expected := *mockedFlagUtil("sec-key", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.secKey)
}

func TestGetVarsShouldInvokeFlagUtilForGoogleClientId(t *testing.T) {
	expected := *mockedFlagUtil("google-client-id", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.googleProvider.clientId)
}

func TestGetVarsShouldInvokeFlagUtilForGoogleSecret(t *testing.T) {
	expected := *mockedFlagUtil("google-secret", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.googleProvider.clientSecret)
}

func TestGetVarsShouldInvokeFlagUtilForGoogleRedirectUrl(t *testing.T) {
	expected := *mockedFlagUtil("google-redirect-url", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.googleProvider.redirectUrl)
}

func TestGetVarsShouldInvokeFlagUtilForRedirectUrl(t *testing.T) {
	expected := *mockedFlagUtil("redirect-url", "", "", false)
	vars := GetVars(mockedFlagUtil)
	assert.Equal(t, expected, vars.redirectUrl)
}

func init() {
	log.SetOutput(ioutil.Discard)
}

func beforeTest() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func mockedFlagUtil(name string, fullName string, defaultValue string, mandatory bool) *string {
	value := name + " expected value"
	return &value
}

func mockedFlagUtilWithEmptyReturn(name string, fullName string, defaultValue string, mandatory bool) *string {
	value := ""
	return &value
}

var TestVars = Vars {
	domain: "MY_DOMAIN",
	port: "1234",
	secKey: "MY_SECURITY_KEY",
	googleProvider: provider {
		clientId: "MY_GOOGLE_CLIENT_ID",
		clientSecret: "MY_GOOGLE_CLIENT_SECRET",
		redirectUrl: "MY_REDIRECT_URL_AFTER_GOOGLE",
	},
}
