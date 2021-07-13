# Infura API Challenge

This app exposes Ethereum Mainnet account, transaction, and block data from the INFURA JSON-RPC API via REST endpoints. The app relies on `gorilla/mux`, a powerful URL router and dispatcher and `gorilal/handlers`, a collection of useful handlers for Go's net/http package.

## Directory Structure

```
infra-test-ethan-blumenthal
    |- app/                     - Contains main API logic files
        |- handler.go           - Defines handler methods of endpoints
        |- router.go            - Defines routes and endpoints
        |- benchmark_test.go    - Defines server load tests
    |- Dockerfile               - Dockerfile for Docker container
    |- docker-compose.yml       - Defines multi-container environment
    |- README.md                - Readme file for documentation
    |- main.go                  - Entry point of the API
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

#### Running locally

Make sure to set `PROJECT_ID` and `PORT` as environment variables.

```
go run main.go
```

#### Running with Docker

```
docker-compose up -d
```

## Performance

Run benchmark tests on the server at various loads.

```
GOMAXPROCS=1 go test -bench=NetHTTPServerGet -benchmem -benchtime=10s
```

## TO DO

```
    |- Complete setup of testing server
    |- Complete cache integration
```
