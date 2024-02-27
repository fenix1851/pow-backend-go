# PoW Quotes Server

The PoW Quotes Server is a Go-based microservice that serves quotes while employing a Proof of Work (PoW) system to prevent abuse. It uses the Gin framework for routing HTTP requests.

## Features

- **Proof of Work (PoW) Challenge:** Clients must solve a PoW challenge to retrieve quotes, preventing spam by requiring computational work for each request.
- **Adaptive Difficulty:** The PoW challenge difficulty adjusts based on request volume within a set timeframe.
- **Quotes Database Management:** Capable of loading quotes from a JSON file into a persistent database.

## Endpoints

- `GET /api/v1/pow/`: Returns data for the PoW challenge, including a random number, a token, and the leading zeros count.
- `GET /api/v1/quotes/`: Retrieves quotes from the database, requiring a valid PoW token.

## Running the Server

### Without Docker

Ensure Go is installed and perform the following:

1. Change to the server's directory.
2. Install dependencies with `go mod tidy`.
3. Launch the server with `go run main.go` and the necessary flags.

### With Docker

Just run command:

```shell
docker compose up --build
```
in a root directory of the project.

## Configuration Flags

```shell
-leadingZerosUpdateInterval int  # Update interval for leading zeros (default 1)
-port string                     # Server port (default "8080")
-quotesJsonPath string           # Path to quotes JSON file
-requestsThreshold int           # Request threshold for PoW calculation (default 1)
-timeFrame int                   # Time frame for requests count (default 20)
```

# PoW Quotes Client

The client interacts with the PoW Quotes Server, solving the PoW challenge and fetching quotes using the obtained token.

## Configuration Flag

```shell
-baseUrl string  # Base URL of the quotes server (default "http://localhost:8080")
```

### Running the Client

#### Without Docker

With Go installed:

1. Change to the client's directory.
2. Install dependencies.
3. Execute `go run main.go` with the server's base URL as a flag.

#### With Docker

Just run command:

```shell
docker compose up --build
```

In a root directory of the project.

Than you can use the client inside the container.

 ```shell
docker exec -it <container_id> /bin/sh
```

Or you still can build the client and run it on your local machine, because server port is exposed.

## Docker Compose Example

Below is a `docker-compose.yml` snippet for running the server:

```yaml
version: '3'

services:
  quotes-server:
    build: ./quotes-server # path to server Dockerfile
    ports:
      - "8080:8080"
    volumes:
      - ./quotes-server/quotes.json:/app/quotes.json # Volume for quotes.json file
    command: >
      sh -c "go run main.go
      -leadingZerosUpdateInterval 1
      -port 8080
      -quotesJsonPath /app/quotes.json
      -requestsThreshold 1
      -timeFrame 20"
    depends_on:
      - quotes-client

  quotes-client:
    depends_on:
      - quotes-server
    build: ./quotes-client # path to client Dockerfile
    command: >
      sh -c "go run main.go
      -baseUrl http://quotes-server:8080"

```

This `docker-compose` setup ensures that the server and client are deployed together and can communicate effectively.
