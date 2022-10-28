package cmd

import (
	"fmt"
	"os"
)

<<<<<<< Updated upstream
// rootCmd represents the base command when called without any subcommands
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
=======
var (
	persistRootFlag bool
	rootCmd         = &cobra.Command{
		Use:   "tweeterapp",
		Short: "A brief description of your application",
		Long:  `A longer description `,
	}
)
var searchCmd = &cobra.Command{
	Use:   "searching for twitter handles",
	Short: "",
	Long:  "",
}
>>>>>>> Stashed changes

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
<<<<<<< Updated upstream
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")
	rootCmd.AddCommand(searchCmd)
=======
	rootCmd.PersistentFlags().BoolvarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")
>>>>>>> Stashed changes
}
