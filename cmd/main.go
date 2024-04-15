package main

import (
	"os"

	"github.com/gardusig/guessclient/guess"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Debug("starting to create guess client...")
	numberGuessClient, err := guess.NewGuessClient()
	if err != nil {
		logrus.Error("failed to create guess client, reason:", err)
		os.Exit(1)
	}
	logrus.Debug("done creating guess client")
	openedBox, err := numberGuessClient.GetBox()
	if err != nil {
		logrus.Error("failed to get opened box, reason:", err)
		os.Exit(1)
	}
	if openedBox == nil {
		logrus.Error("expected opened box, got nil instead")
		os.Exit(1)
	}
	logrus.Debug("message: ", openedBox.Message)
}
