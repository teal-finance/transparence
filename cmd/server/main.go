package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	addr := flag.String("p", "8080", "listening port")
	flag.Parse()

	s := &Service{}

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", *addr),
		Handler: s.Handler(),
	}

	log.Println("Server listening on", server.Addr)
	log.Fatal(server.ListenAndServe())
}

type Service struct{}

func (s *Service) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middlewareLog)
	r.Get("/{id}", s.handleMeasurement)
	return r
}

func (s *Service) handleMeasurement(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// write here ...

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func middlewareLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
