package cmd

import (
	"fmt"

	"git.dwarvesf.com/url-shortener/urlshorten/db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a urlshorten from the list",
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("delete")
		err := db.DeleteURLShorten(key)
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		fmt.Println("Delete Success")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("delete", "d", "", "Remove a urlshorten from the list")
}
