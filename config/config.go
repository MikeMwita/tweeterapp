package config

import (
	"errors"
	"os"
)

type Config struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func NewConfig() (*Config, error) {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	ck, ok := os.LookupEnv("Twconsumerkey")
=======
	ck, ok := os.LookupEnv("CONSUMER_KEY")
>>>>>>> Stashed changes
=======
	ck, ok := os.LookupEnv("CONSUMER_KEY")
>>>>>>> Stashed changes
	if !ok {
		return nil, errors.New("consumer key need to be set")
	}
	csk, ok := os.LookupEnv("CONSUMER_SECRET")
	if !ok {
		return nil, errors.New(" Consumer_secret key needs to be set")
	}
	At, ok := os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		return nil, errors.New("the access_token is missing")
	}

	Ats, ok := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !ok {
		return nil, errors.New("the access_token secret key is missing")
	}
	return &Config{
		ConsumerKey:       ck,
		ConsumerSecret:    csk,
		AccessToken:       At,
		AccessTokenSecret: Ats,
	}, nil
}
