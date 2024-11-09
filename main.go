package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}

func (s *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create users"))
}

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()
	srv := &http.Server{Addr: api.addr, Handler: mux}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv.ListenAndServe()
}
