package cmd

import (
	"fmt"

	"git.dwarvesf.com/url-shortener/urlshorten/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all urlshorten in DB.",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := db.AllURLShorten()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}

		if len(list) == 0 {
			fmt.Println("You don't have any urlshorten")
		}

		for _, value := range list {
			fmt.Printf("%s - %s", value.Key, value.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("list", "l", "", "List all urlshorten in DB")
}
