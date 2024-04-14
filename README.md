# GuessClient

## Project Overview

GuessClient is the client-side application for the Guess service. It sends guess requests to a server using gRPC.

## How to Build and Run

To build and run GuessClient locally, follow these steps:

```bash
$ docker build . -t guessclient
$ docker run guessclient
```

## Configuration Options
GuessClient can be configured using environment variables:

- `GUESS_SERVER_GRPC_PORT`: The **port** of the Guess server gRPC service (default: `50051`).
- `GUESS_SERVER_GRPC_HOST`: The **host** of the Guess server gRPC service (default: `localhost`).

## Code Structure
The project directory structure is as follows:

```
guessclient/
├── cmd/
├── guesser/
└── internal/
```

- `cmd/`: Contains the main entry point of the application.
- `guesser/`: Contains the client logic for interacting with the Guess service.
- `internal/`: Contains configuration constants and initialization logic.
