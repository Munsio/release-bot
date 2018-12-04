package server

import (
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// Server struct
type Server struct {
	port   int
	router *http.ServeMux
}

// NewServer creates the http listener
func NewServer(port int) *Server {
	mux := http.NewServeMux()
	return &Server{
		port:   port,
		router: mux,
	}
}

// AddRoute ads an handler with path to the muxer
func (serv *Server) AddRoute(path string, handle http.Handler) {
	serv.router.Handle(path, handle)
}

// Run the server
func (serv *Server) Run() {
	log.Info("listen on port: " + strconv.Itoa(serv.port))
	http.ListenAndServe(":"+strconv.Itoa(serv.port), serv.router)
}
