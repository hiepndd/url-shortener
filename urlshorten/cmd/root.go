package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is root save all sub-command
var RootCmd = &cobra.Command{
	Use:   "urlshorten",
	Short: "UrlShorten is an http.Handler that forwards paths to other URLs",
}
