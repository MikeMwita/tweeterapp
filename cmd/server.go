package cmd

import (
	"fmt"
	"log"
)
var (
	localRootFlag bool
serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Used to listen and serve the tweets",
	Long:  `Listens for tweets with the hashtag Kipchoge and serves them via twitterAPI:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
)
}
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
func init() {
	fmt.Println("Inside the init function")
	cobra.onInitialize(initConfig)
	rootCmd.AddCommand(serverCmd)
	rootCmd.PersistentFlags().BoolvarP(&localRootFlag,"local flag","l",false,"local root flag")
}
