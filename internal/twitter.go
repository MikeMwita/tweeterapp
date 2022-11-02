package internal

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/MikeMwita/tweeterapp/internal/common"
	"github.com/MikeMwita/tweeterapp/internal/platform/twitter"
)

func NewClient(config *config.Config) (common.TweetInterface, error) {
	return twitter.NewClient(config)
}
