package twitter

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/MikeMwita/tweeterapp/internal/common"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)
type Client struct {

}

var Client = (common.SocialInterface())(nil)

func NewClient(cfg *config.Config) (*Tw, error) {
	config := oauth1.NewConfig(config.ck, config.csk)
	token := oauth1.NewToken(config.At, config.Ats)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	client,err :=getClient(&config){
		if err!=nil{
			log.Println("The tweeter client is unreachable")
			log.Println(err)
		}
	}

	tweet, resp, err := client.Statuses.Update("setting up a twitter client", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n",tweet)
	log.Printf("%v\n",resp)
}

type search struct {
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:"kipchoge"
	})

	if err!=nil{
		log.Fatal(err)
}
	log.printf("%v",search)
	log.printf("%v",resp)

}

