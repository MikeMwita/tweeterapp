package config

import (
	"errors"
	"os"
)

type Config struct {
	TwConsumerKey     string
	TwConsumerSecret  string
	AccessToken       string
	AccessTokenSecret string
}

func NewConfig() (*Config, error) {
	ck, ok := os.LookupEnv("Twconsumerkey")
	if !ok {
		return nil, errors.New("CK need to be set")
	}

	csk, ok := os.LookupEnv("TwConsumerSecret")
	if !ok {
		return nil, errors.New("The consumer secret key needs to be set")
	}

	At, ok := os.LookupEnv("AccessToken")
	if !ok {
		return nil, errors.New("the access token is missing")
	}

	Ats, ok := os.LookupEnv("AccessTokenSecret")
	if !ok {
		return nil, errors.New("the accesstoken secret key is missing")
	}

	return &Config{
		TwConsumerKey:     ck,
		TwConsumerSecret:  csk,
		AccessToken:       At,
		AccessTokenSecret: Ats,
	}, nil
}
