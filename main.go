package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	s := &server{addr: ":8080"}
	err := http.ListenAndServe(s.addr, s)
	if err != nil {
		log.Fatal(err)
	}
}
