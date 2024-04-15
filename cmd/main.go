package main

import (
	"os"

	"github.com/gardusig/guessclient/guesser"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Debug("starting to create number guesser client...")
	numberGuesser, err := guesser.NewGuesser()
	if err != nil {
		logrus.Error("failed to create number guesser, reason:", err)
		os.Exit(1)
	}
	logrus.Debug("done creating number guesser")
	openedBox, err := numberGuesser.GetBox()
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
