/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	"github.com/MikeMwita/tweeterapp.git/cmd"
=======
	"github.com/MikeMwita/tweeterapp/cmd"
	"github.com/MikeMwita/tweeterapp/internal/platform/twitter"
	"log"
>>>>>>> Stashed changes
=======
	"github.com/MikeMwita/tweeterapp/cmd"
	"github.com/MikeMwita/tweeterapp/internal/platform/twitter"
	"log"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	fmt.Println("Our twitter bot")

=======
	getClient()
	fmt.Println("Our twitter bot")
>>>>>>> Stashed changes
=======
	getClient()
	fmt.Println("Our twitter bot")
>>>>>>> Stashed changes
	cmd.Execute()

}
