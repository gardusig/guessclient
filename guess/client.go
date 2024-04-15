package guess

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/gardusig/guessclient/internal"
	guessproto "github.com/gardusig/guessproto/generated/go"
)

const maxRetryAttempt = 5

type GuessServiceClient struct {
	guessproto.GuessServiceClient

	connection *grpc.ClientConn
}

func NewGuessServiceClient() (*GuessServiceClient, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", internal.GuessServiceHost, internal.GuessServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := guessproto.NewGuessServiceClient(conn)
	return &GuessServiceClient{
		GuessServiceClient: client,
		connection:         conn,
	}, nil
}

func (c *GuessServiceClient) SendGuessRequest(level uint32, guess int64) (*guessproto.GuessNumberResponse, error) {
	request := guessproto.GuessNumberRequest{
		Level: level,
		Guess: guess,
	}
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := c.GuessNumber(context.Background(), &request)
		if err == nil {
			return resp, nil
		}
		_, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("failed to make guess request after %v attempts", maxRetryAttempt)
}

func (c *GuessServiceClient) SendOpenBoxRequest(lockedBox *guessproto.LockedBox) (*guessproto.OpenedBox, error) {
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := c.OpenBox(context.Background(), lockedBox)
		if err == nil {
			return resp, nil
		}
		_, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("failed to make openBox request after %v attempts", maxRetryAttempt)
}

func (c *GuessServiceClient) CloseConnection() {
	c.connection.Close()
}
