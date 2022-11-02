package twitter

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/MikeMwita/tweeterapp/internal/common"
	"github.com/MikeMwita/tweeterapp/internal/models"
	"net/http"
)

type Client struct {
	tweet *http.Client
}

func (c Client) Auth() error {
	//TODO implement me
	panic("implement me")
}

func (c Client) Search() ([]models.Tweet, error) {
	//TODO implement me
	panic("implement me")
}

func NewClient(cfg *config.Config) (common.TweetInterface, error) {
	return &Client{}, nil
}
