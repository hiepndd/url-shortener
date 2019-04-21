package cmd

import (
	"fmt"
	"net/http"

	"git.dwarvesf.com/url-shortener/urlshorten/db"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run HTTP server on a given port",
	Run: func(cmd *cobra.Command, args []string) {
		run, _ := cmd.Flags().GetString("port")
		fmt.Println(run)
		mux := defaultMux()
		maps := buildMap()
		handler := mapHandler(maps, mux)
		newURL := ":" + run
		fmt.Printf("Starting the server on %s", newURL)
		http.ListenAndServe(newURL, handler)

	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("port", "p", "", "Run HTTP server on a given port")
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloworld)
	return mux
}

func buildMap() map[string]string {
	pathToUrls := make(map[string]string)
	list, _ := db.AllURLShorten()
	for _, path := range list {
		pathToUrls[path.Key] = path.Value
	}
	return pathToUrls

}

func mapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			err := db.Count(pathToUrls[path])
			if err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
