package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "configure -a",
	Short: "Adds a urlshorten to your url list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
