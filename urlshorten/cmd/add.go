package cmd

import (
	"fmt"

	"git.dwarvesf.com/url-shortener/urlshorten/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "configure",
	Short: "Adds a urlshorten to your url list.",
	Run: func(cmd *cobra.Command, args []string) {
		append, _ := cmd.Flags().GetString("append")
		url, _ := cmd.Flags().GetString("url")
		err := db.AddURLToYamlFile(append, url)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Println("Added to your DB")

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("append", "a", "", "Append a url to list")
	addCmd.Flags().StringP("url", "u", "", "URL")
}
