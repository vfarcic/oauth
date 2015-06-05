package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"log"
)

type Vars struct {
	domain string
	port string
	secKey string
	redirectUrl string
	googleProvider provider
	facebookProvider provider
	// TODO: Add the rest of providers
	//	githubProvider provider
	//	herokuProvider provider
	//	soundcloudProvider provider
}

type provider struct {
	clientId string
	clientSecret string
	redirectUrl string
}

func GetVars(flagUtilFunc flagUtilGetter) Vars {
	domain := flagUtilFunc("domain", "Application domain", "", false)
	port := flagUtilFunc("port", "Application port", "8080", false)
	secKey := flagUtilFunc("sec-key", "Security key", "", true)
	redirectUrl := flagUtilFunc("redirect-url", "Redirect URL", "", false)
	googleClientId := flagUtilFunc("google-client-id", "Google Client ID", "", false)
	googleSecret := flagUtilFunc("google-secret", "Google Secret", "", false)
	googleRedirectUrl := flagUtilFunc("google-redirect-url", "Google Redirect URL", "", false)
	flag.Parse()
	googleProvider := getProvider("Google", *googleClientId, *googleSecret, *googleRedirectUrl)
	if len(*secKey) == 0 {
		fmt.Println("Security key is mandatory")
		os.Exit(1)
	}
	if len(*redirectUrl) == 0 {
		fmt.Println("Redirect URL is mandatory")
		os.Exit(1)
	}
	log.Println("Google OAuth is set")
	return Vars{
		domain: *domain,
		port: *port,
		secKey: *secKey,
		redirectUrl: *redirectUrl,
		googleProvider: *googleProvider,
	}
}

func getProvider(name string, clientId string, secret string, redirectUrl string) *provider {
	if len(clientId) > 0 && (len(secret) == 0 || len(redirectUrl) == 0) {
		fmt.Println("All", name, "data needs to be set when", strings.ToLower(name) + "-client-id", "is specified")
		os.Exit(1)
	}
	return &provider{
		clientId,
		secret,
		redirectUrl,
	}
}

func getFlagArgs(
	name string,
	fullName string,
	defaultValue string,
	mandatory bool) (argName string, argValue string, argUsage string) {
	envName := strings.Replace(strings.ToUpper(name), "-", "_", -1)
	usageMandatory := ""
	if mandatory {
		usageMandatory = "This variable is mandatory."
	}
	value := os.Getenv(envName)
	if len(value) == 0 {
		value = defaultValue
	}
	usage := fullName + ". " +
			usageMandatory +
			"It can also be specified with environment variable " + envName + "."
	return name, value, usage
}

type flagUtilGetter func(name string, fullName string, defaultValue string, mandatory bool) *string

func flagUtil(name string, fullName string, defaultValue string, mandatory bool) *string {
	name, value, usage := getFlagArgs(name, fullName, defaultValue, mandatory)
	return flag.String(name, value, usage)
}
