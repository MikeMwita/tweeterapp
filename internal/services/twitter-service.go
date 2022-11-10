package services

import "github.com/MikeMwita/tweeterapp/internal/common"

type Tweet interface{
searchTweet() error
LikeTweet() error
RetweetTweet()error

//search ,like --a tweet based on the ID
searchTweetById(tweetId int)([]*Tweet,error)
LikeweetById(tweetId int)([]*Tweet,error)
RetweetTweetById(tweetId int)([]*Tweet,error)

}

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
