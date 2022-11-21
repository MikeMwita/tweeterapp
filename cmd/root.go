package cmd

import (
	"fmt"
	"os"
)

var (
	persistRootFlag bool
	rootCmd         = &cobra.Command{
		Use:   "tweeterapp.git",
		Short: "A tweeter robot application ",
		Long:  `This is a simple application that perfoms various functions  through twitter  platform:`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the root command ")
		},
	}
	searchCmd = &cobra.Command{
		Use:   "for searching through the tweets",
		Short: "Look for a tweet with  the name Kipchoge ",
		Long:  "Listens through the twitterAPI to get a tweet with the name Kipchoge ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tweet")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.PersistentFlags().BoolvarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")

}
