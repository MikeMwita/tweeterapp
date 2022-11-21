package cmd

import (
	"fmt"

<<<<<<< Updated upstream
	"github.com/spf13/cobra"
)


var searchCmd=&cobra.Command{
Use :"search through the tweets",
Short: "",
Long: "",
Run: func(cmd *cobra.Command,args[]string){
	fmt.Println("Searching for the tweets")
	err:=app.Initialize()
	if err!=nil{
panic(err)
	}
},

}

func init(){

rootCmd.AddCommand(searchCmd)

}
=======
	"github.com/MikeMwita/tweeterapp/internal/app"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{

Use:"Used for searching the twitter handles ",
Short: "Searches for the twiiter handles that have the name Kipchoge in them",
Long:"",

Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("Searching for the tweets..")
	err := app.Initialize()
	if err != nil {
		panic(err)
	}
},



}
>>>>>>> Stashed changes
