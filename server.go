package main

import (
	"infura-challenge/client"
	"infura-challenge/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Args args used to run the server
type Args struct {
	// infura project id of the form,
	// e.g "2df95ac72e5a4153b3de94977e4d3783"
	projectId string
	// port for the server of the form,
	// e.g ":8080"
	port string
}

// Run run the server based on given args
func Run(args Args) error {
	// router
	router := mux.NewRouter().
		PathPrefix("/api/v1/"). // add prefix for v1 api `/api/v1/`
		Subrouter()

	client := client.NewInfuraJSONRPCClient(args.projectId)
	handler := handlers.NewEventHandler(client)
	
	// set content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// get blocks
	router.HandleFunc("/block", handler.Get).Methods(http.MethodGet)
	// get transactions
	router.HandleFunc("/transaction", handler.Get).Methods(http.MethodGet)

	// start server
	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}