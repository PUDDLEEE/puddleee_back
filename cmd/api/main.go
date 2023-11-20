package main

import (
	"net/http"
)

type application struct {
	router http.Handler
}

func main() {
	config := new(application)
	config.router = config.routers()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: config.router,
	}

	srv.ListenAndServe()
}
