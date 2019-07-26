package http

import (
	// stdlib
	"log"
	"net"
	"net/http"
	// universe
	"github.com/universelabs/universe-server/universe"
)

// HTTP service
type Server struct {
	// net/http infrastructure
	ln net.Listener
	// handler to serve
	Handler *Handler
	// bind address to open
	Addr string
}

// Returns a new Server instantiated from the arguments 
func NewServer(addr string, ks universe.Keystore) *Server {
	srv := &Server {
		// set addr
		Addr: addr,
		// init handler
		Handler: NewHandler(ks),
	}
	return srv
}

// Listens and serves the server instance
func (s *Server) Open() error {
	// open the socket
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	// start HTTP server
	go func() { log.Fatal(http.Serve(s.ln, s.Handler)) }()
	// *** because http.Serve is called in a go routine, main() must hang
	// the process so that the server doesn't close!

	return nil
}

// Closes the server instance
func (s *Server) Close() error {
	if s.ln != nil {
		return s.ln.Close()
	}
	return nil
}