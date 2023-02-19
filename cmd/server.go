package cmd

import (
	"fmt"
	"jag-cli/internal/generate"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts an HTTP server",
	RunE:  serverHandler,
}

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "the port to listen on")
	rootCmd.AddCommand(serverCmd)
}

func serverHandler(cmd *cobra.Command, args []string) error {
	http.HandleFunc("/payload/", func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("received request")

		sizeStr := strings.TrimPrefix(r.URL.Path, "/payload/")
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			http.Error(w, "Invalid size parameter", http.StatusBadRequest)
			return
		}

		payload, err := generate.JsonPayload(size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("received request")

		fmt.Fprintln(w, "Acknowledged!")
	})

	addr := fmt.Sprintf(":%d", port)
	logrus.Infof("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(addr, nil))
	return nil
}
