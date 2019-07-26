package http_test

import (
	// stdlib
	"testing"
	// universe
	"github.com/universelabs/universe-server/http"
)

// test wrapper for http.Server
type Server struct {
	*http.Server
	// test wrapper for Handler
	Handler *Handler
}

// returns a new instance of Server
func NewServer() *Server {
	// create server wrapper
	s := &Server{
		Server: http.NewServer(), 
		Handler: NewHandler(),
	}
	s.Server.Handler = s.Handler.Handler
	
	// use random port
	s.Addr = ":0" 

	return s
}

// returns a running server, panics on error
func MustOpenServer() *Server {
	s := NewServer()
	if err := s.Open(); err != nil {
		panic(err)
	}
	return s
}