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
		list, err := db.AllURLYamlFile()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}

		if len(list) == 0 {
			fmt.Println("You don't have any urlshorten")
		}

		for key, value := range list {
			fmt.Printf("%s - %s\n", key, value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("list", "l", "", "List all urlshorten in DB")
}
