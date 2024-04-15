package guess

import (
	"fmt"

	"github.com/gardusig/guessclient/internal"
	guessproto "github.com/gardusig/guessproto/generated/go"
	"github.com/sirupsen/logrus"
)

type GuessClient struct {
	serviceClient *GuessServiceClient

	level      uint32
	lowerBound int64
	upperBound int64
}

func NewGuessClient() (*GuessClient, error) {
	serviceClient, err := NewGuessServiceClient()
	if err != nil {
		return nil, err
	}
	return &GuessClient{
		serviceClient: serviceClient,
	}, nil
}

func (g *GuessClient) GetBox() (*guessproto.OpenedBox, error) {
	var lockedBox *guessproto.LockedBox
	var err error
	g.level = internal.LevelMinThreshold
	for g.level <= internal.LevelMaxThreshold {
		lockedBox, err = g.guessNumberByLevel()
		if err != nil {
			return nil, err
		}
		if lockedBox != nil {
			g.level += 1
			logrus.Debug("Passed to level: ", g.level, ", encryptedMessage: ", lockedBox.EncryptedMessage)
		}
	}
	return g.serviceClient.SendOpenBoxRequest(lockedBox)
}

func (g *GuessClient) guessNumberByLevel() (*guessproto.LockedBox, error) {
	logrus.Debug("attempt to guess number for level: ", g.level)
	g.lowerBound = internal.GuessMinThreshold
	g.upperBound = internal.GuessMaxThreshold
	for g.lowerBound <= g.upperBound {
		guess := g.lowerBound + ((g.upperBound - g.lowerBound) >> 1)
		logrus.Debug("lowerBound:", g.lowerBound, ", upperBound:", g.upperBound, ", guess:", guess)
		resp, err := g.serviceClient.SendGuessRequest(g.level, guess)
		if err != nil {
			return nil, err
		}
		logrus.Debug("server response:", resp.Result)
		switch resp.Result {
		case internal.Equal:
			return resp.LockedBox, nil
		case internal.Greater:
			g.upperBound = guess - 1
		case internal.Less:
			g.lowerBound = guess + 1
		default:
			return nil, fmt.Errorf("unexpected response from server: %v", resp.Result)
		}
	}
	return nil, fmt.Errorf("failed to guess the right number :/")
}
