package twitter

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/dghubble/go-twitter/twitter"
)

type Tw struct {
	client
}

func NewClient(cfg *config.Config) (*Tw, error) {
	httpAuth, err := generateAuth(cfg)
	c := twitter.NewClient()
}

func generateAuth(cfg *config.Config) {
	cfg := oauth1
}
