package main

import (
	"fmt"
	"os"

	"git.dwarvesf.com/url-shortener/urlshorten/cmd"
	"git.dwarvesf.com/url-shortener/urlshorten/db"
)

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	must(db.Init())
	must(cmd.RootCmd.Execute())
}
