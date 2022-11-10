package cmd

import (
	"fmt"

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