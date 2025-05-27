package main

import (
	"fullGen/internal/handlers"
	tt "fullGen/tintin"
	"net/http"
)

func CreateAndStartServer() {
	server := handlers.NewServer()
	newServer, err := tt.NewServer(server)
	if err != nil {
		panic("failed to create new server") // panic if some of must- functions are failed
	}
	err = http.ListenAndServe(":8080", newServer)
	if err != nil {
		panic("failed to start server")
	}
}
