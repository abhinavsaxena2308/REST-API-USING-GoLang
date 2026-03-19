package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/go-chi/cors"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	// creating router
	router := mux.NewRouter()

	//register services
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:5173"},
	// 	AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders: []string{"Content-Type", "Authorization"},
	// 	// AllowedCredentials: true,
	// })
	// corsHandler := c.Handler(router)

	log.Println("listening on port ", s.addr)

	return http.ListenAndServe(s.addr, router)
}