package main

import (
	"log"
	"os"

	"github.com/augsync/teamcore-project-46344/internal/httptransport"
	"github.com/augsync/teamcore-project-46344/internal/service"
	"github.com/valyala/fasthttp"
)

func main() {
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	s := service.New()

	log.Printf("listening on port %s", port)
	log.Fatal(fasthttp.ListenAndServe(":"+port, httptransport.Handler(s).Handler))
}
