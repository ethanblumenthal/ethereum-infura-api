package app

import (
	"log"
	"net/http"

	"github.com/INFURA/go-libs/jsonrpc_client"
	"github.com/coocood/freecache"
	"github.com/gorilla/mux"
)

// In bytes, where 1024 * 1024 represents a single Megabyte, and 100 * 1024*1024 represents 100 Megabytes.
var cacheSize = 100 * 1024 * 1024 * 10000000000
var handler = &Handler{EthereumClient: jsonrpc_client.EthereumClient{URL: "https://mainnet.infura.io/v3/2df95ac72e5a4153b3de94977e4d3783"}, Cache: freecache.Cache{}}

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
		"/blocks/number",
		handler.GetBlockNumber,
	},
	Route {
		"GetBlockByNumber",
		"GET",
		"/blocks/number/{num}",
		handler.GetBlockByNumber,
	},
	Route {
		"GetBlockByHash",
		"GET",
		"/blocks/{hash}",
		handler.GetBlockByHash,
	},
    Route {
		"GetTransactionByHash",
		"GET",
		"/transactions/{hash}",
		handler.GetTransactionByHash,
	},
}

// NewRouter configures a new router for the API
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