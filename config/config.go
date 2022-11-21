package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
<<<<<<< Updated upstream
	ConsumerKey       string
	ConsumerSecret    string
=======
	ConsumerKey     string
	ConsumerSecret  string
>>>>>>> Stashed changes
	AccessToken       string
	AccessTokenSecret string
}

func NewConfig() (*Config, error) {
<<<<<<< Updated upstream
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	ck, ok := os.LookupEnv("Twconsumerkey")
=======
	ck, ok := os.LookupEnv("CONSUMER_KEY")
>>>>>>> Stashed changes
=======
	ck, ok := os.LookupEnv("CONSUMER_KEY")
>>>>>>> Stashed changes
=======

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	  }

	ck, ok := os.LookupEnv("CONSUMER_KEY")
>>>>>>> Stashed changes
	if !ok {
		return nil, errors.New("consumer key need to be set")
	}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	csk, ok := os.LookupEnv("CONSUMER_SECRET")
	if !ok {
		return nil, errors.New(" Consumer_secret key needs to be set")
	}
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	At, ok := os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		return nil, errors.New("the access_token is missing")
	}

	Ats, ok := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !ok {
		return nil, errors.New("the access_token secret key is missing")
	}
	return &Config{
<<<<<<< Updated upstream
		ConsumerKey:       ck,
		ConsumerSecret:    csk,
=======
		ConsumerKey:     ck,
		ConsumerSecret:  csk,
>>>>>>> Stashed changes
		AccessToken:       At,
		AccessTokenSecret: Ats,
	}, nil
}
