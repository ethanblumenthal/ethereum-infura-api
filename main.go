package main

import (
	"log"
	"os"
)

func main() {
	args := Args{
		projectId: "2df95ac72e5a4153b3de94977e4d3783",
		port: ":8080",
	}
	if projectId := os.Getenv("PROJECT_ID"); projectId != "" {
		args.projectId = projectId
	}
	if port := os.Getenv("PORT"); port != "" {
		args.port = ":" + port
	}
	// run server
	if err := Run(args); err != nil {
		log.Println(err)
	}
}