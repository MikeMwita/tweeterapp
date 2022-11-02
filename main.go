package main

import (
	"fmt"
	"github.com/MikeMwita/tweeterapp.git/cmd"
	"github.com/MikeMwita/tweeterapp.git/internal/platform/twitter"
	"log"
)

func getClient() error {
	twitter.NewClient(httpClient)
	client, err := getClient(&config)
	if err != nil {
		log.Println("The tweeter client is unreachable")
		log.Println(err)
	}
}
func main() {

	getClient()
	fmt.Println("Our twitter bot")
	cmd.Execute()

}
