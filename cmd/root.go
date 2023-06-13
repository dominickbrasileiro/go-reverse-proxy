package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var (
	origin string
	host   string
	port   string
)

var rootCmd = &cobra.Command{
	Use:   "go-reverse-proxy",
	Short: "Redirect requests to another server easily",
	Long:  `Redirect requests to another server easily, without the need to change the DNS or the port of the origin server.`,
	Run: func(cmd *cobra.Command, args []string) {
		originServerURL, err := url.Parse(origin)
		if err != nil {
			log.Fatal("invalid origin server URL")
		}

		reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			req.Host = originServerURL.Host
			req.URL.Host = originServerURL.Host
			req.URL.Scheme = originServerURL.Scheme
			req.RequestURI = ""

			originServerResponse, err := http.DefaultClient.Do(req)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(rw, err)
				return
			}

			rw.WriteHeader(http.StatusOK)
			io.Copy(rw, originServerResponse.Body)
		})

		log.Fatal(http.ListenAndServe(host+":"+port, reverseProxy))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&origin, "origin", "o", "", "Origin server URL")
	rootCmd.Flags().StringVarP(&host, "host", "H", "0.0.0.0", "Host to listen to")
	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen to")

	err := rootCmd.MarkFlagRequired("origin")
	if err != nil {
		log.Fatal(err)
	}
}
