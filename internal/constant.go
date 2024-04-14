package internal

import "os"

const (
	LevelMinThreshold uint32 = 0
	LevelMaxThreshold uint32 = 1000

	GuessMinThreshold int64 = -4000000000000000000
	GuessMaxThreshold int64 = +4000000000000000000

	guessServicePortEnvKey = "GUESS_SERVER_GRPC_PORT"
	guessServiceHostEnvKey = "GUESS_SERVER_GRPC_HOST"

	guessServiceDefaultPort = "50051"
	guessServiceDefaultHost = "localhost"
)

var (
	Equal   = "="
	Less    = "<"
	Greater = ">"

	GuessServicePort string
	GuessServiceHost string
)

func init() {
	GuessServicePort = os.Getenv(guessServicePortEnvKey)
	if GuessServicePort == "" {
		GuessServicePort = guessServiceDefaultPort
	}
	GuessServiceHost = os.Getenv(guessServiceHostEnvKey)
	if GuessServiceHost == "" {
		GuessServiceHost = guessServiceDefaultHost
	}
}
