package main

import (
	"fmt"
	"os"
	"path/filepath"

	"git.dwarvesf.com/url-shortener/urlshorten/cmd"
	"git.dwarvesf.com/url-shortener/urlshorten/db"
	"github.com/mitchellh/go-homedir"
)

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "urlshorten.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}
