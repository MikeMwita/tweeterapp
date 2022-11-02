package common

import "github.com/MikeMwita/tweeterapp/internal/models"

type TweetInterface interface {
	Auth() error
	Search() ([]models.Tweet, error)
}
