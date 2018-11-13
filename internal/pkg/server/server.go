package server

import (
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/karriereat/release-bot/internal/pkg/config"
)

// Server struct
type Server struct {
	Config *config.Config
	router *http.ServeMux
}

// NewServer creates the http listener
func NewServer(conf *config.Config) *Server {
	mux := http.NewServeMux()
	return &Server{
		Config: conf,
		router: mux,
	}
}

// AddRoute ads an handler with path to the muxer
func (serv *Server) AddRoute(path string, handle http.Handler) {
	serv.router.Handle(path, handle)
}

// Run the server
func (serv *Server) Run() {
	log.Info("listen on port: " + strconv.Itoa(serv.Config.Port))
	http.ListenAndServe(":"+strconv.Itoa(serv.Config.Port), serv.router)
}
