package webrezpro_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	webrezpro "github.com/omniboost/go-webrezpro"
)

var (
	client *webrezpro.Client
)

func TestMain(m *testing.M) {
	var baseURL *url.URL
	var err error

	baseURLString := os.Getenv("WEBREZPRO_BASE_URL")
	username := os.Getenv("WEBREZPRO_USERNAME")
	password := os.Getenv("WEBREZPRO_PASSWORD")

	client = webrezpro.NewClient(nil)
	client.SetUsername(username)
	client.SetPassword(password)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
