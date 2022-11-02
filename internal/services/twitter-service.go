package services

import "github.com/MikeMwita/tweeterapp/internal/common"

type TwitterService interface {
	Auth() error
}

type DefaultTwService struct {
	tweetClient common.TweetInterface
}

func (d DefaultTwService) Auth() error {
	err := d.tweetClient.Auth()
	if err != nil {
		return err
	}

	return nil
}

func NewDefaultTwService(client common.TweetInterface) TwitterService {
	return &DefaultTwService{tweetClient: client}
}
