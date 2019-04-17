package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "urlshorten",
	Short: "UrlShorten is an http.Handler that forwards paths to other URLs",
}
