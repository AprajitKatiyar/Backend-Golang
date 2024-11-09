package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("Index page"))
			return
		case "/about":
			w.Write([]byte("About page"))
			return
		default:
			w.Write([]byte("404 page not found"))
			return

		}
	}
}
func main() {
	api := &api{addr: ":8080"}
	srv := &http.Server{Addr: api.addr, Handler: api}
	srv.ListenAndServe()
}
