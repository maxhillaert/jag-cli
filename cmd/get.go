package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Perform an HTTP GET request to the specified URL",
	Args:  cobra.ExactArgs(1),
	RunE:  getHandler,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getHandler(cmd *cobra.Command, args []string) error {
	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body:")
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
