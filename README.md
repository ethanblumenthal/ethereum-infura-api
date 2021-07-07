# Infura API Challenge

This app exposes Ethereum Mainnet account, transaction, and block data from the INFURA JSON-RPC API via REST endpoints. The app relies on `gorilla/mux`, a powerful URL router and dispatcher and `gorilal/handlers`, a collection of useful handlers for Go's net/http package.

## Directory Structure

```
infra-test-ethan-blumenthal
    |- app/               - Contains main API logic files
        |- controller.go  - Defines handler methods of endpoints
        |- models.go      - Block and Transaction models
        |- client.go      - Methods interacting with JSON-RPC API
        |- router.go      - Defines routes and endpoints
    |- Dockerfile         - Dockerfile for Docker container
    |- README.md          - Readme file for documentation
    |- main.go            - Entry point of the API
```

## Instructions

#### Install Go 1.16 or higher

Follow the official docs or use your favorite dependency manager
to install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

Verify your `$GOPATH` is correctly set before continuing!

#### Using Go get

```bash
go get -u github.com/INFURA/infra-test-ethan-blumenthal

```

#### Installation

```
go install ./...
```

## Using Docker

Lastly, run docker-compose up and Compose will start and run your entire app.

```
docker-compose up -d
```

## Performance

Run some load test iterations and document the testing approach and the results obtained. Specify some performance expectations given the load test results: e.g., this application is able to support X requests per minute.

## Testing

Run all tests with verbosity but one at a time, without timeout, to avoid ports collisions:

```
go test -v -p=1 -timeout=0 ./...
```
