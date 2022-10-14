package config

import (
	"errors"
	"os"
)

type Config struct {
	TwConsumerKey    string
	TwConsumerSecret string
}

func NewConfig() (*Config, error) {
	ck, ok := os.LookupEnv("TWITTER_CK")
	if !ok {
		return nil, errors.New("CK need to be set")
	}

	return &Config{
		TwConsumerKey:    ck,
		TwConsumerSecret: "",
	}, nil
}
