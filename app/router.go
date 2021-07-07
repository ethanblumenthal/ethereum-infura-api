package app

import (
	"log"
	"net/http"

	"github.com/INFURA/go-libs/jsonrpc_client"
	"github.com/gorilla/mux"
)

var controller = &Controller{EthereumClient: jsonrpc_client.EthereumClient{URL: "https://mainnet.infura.io/v3/2df95ac72e5a4153b3de94977e4d3783"}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
	Route {
		"GetBlockNumber",
		"GET",
		"/blocks",
		controller.GetBlockNumber,
	},
	Route {
		"GetBlockByNumber",
		"GET",
		"/blocks/{num}",
		controller.GetBlockByNumber,
	},
    Route {
		"GetBlockByHash",
		"GET",
		"/blocks/{hash}",
		controller.GetBlockByHash,
	},
    Route {
		"GetTransactionByHash",
		"GET",
		"/transactions/{hash}",
		controller.GetTransactionByHash,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes { 
        var handler http.Handler
		log.Println(route.Name)
        handler = route.HandlerFunc
        
        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}