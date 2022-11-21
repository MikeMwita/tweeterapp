package cmd

import (
	"fmt"
<<<<<<< Updated upstream
=======

	"github.com/MikeMwita/tweeterapp/internal/app"
>>>>>>> Stashed changes
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "This command is used to serve requests",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	erverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
