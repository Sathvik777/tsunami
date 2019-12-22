package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// serveCmd represents the base command when called without any subcommands
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A test server to load test",
	Long:  `A test server to load test and push prometheus metrics`,
	Run:   serveHTTP,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func serveHTTP(cmd *cobra.Command, args []string) {

	http.HandleFunc("/health", health)
	log.Print("Started Server at port ", 8080)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		log.Fatal("Serving exited with error")
	}
}
