package main

import (
	"net/http"
)

func main() {
	api := &api{Addr: ":8080"}
	mux := http.NewServeMux()
	srv := &http.Server{Addr: api.Addr, Handler: mux}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
