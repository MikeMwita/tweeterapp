package twitter

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

type Tw struct {
	Search  *SearchService
	Streams *StreamService
}
func NewClient(cfg *config.Config) (*Tw, error) {
	//passing in the keys
	config := oauth1.NewConfig(config.ck, config.csk)
	//tokens
	token := oauth1.NewToken(config.At, config.Ats)
	httpClient := config.Client(oauth1.NoContext, token)
	
	//initializing the twitter client
	client := twitter.NewClient(httpClient)
	client,err :=getClient(&config){
		if err!=nil{
			log.Println("The tweeter client is unreachable")
			log.Println(err)
		}
	}
	// send a tweet
	tweet, resp, err := client.Statuses.Update("setting up a twitter client", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n",tweet)
	log.Printf("%v\n",resp)
}

//search for a tweet
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


