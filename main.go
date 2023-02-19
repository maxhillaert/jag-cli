package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"jag-cli/cmd"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceQuote:    true,
	})

	logrus.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
