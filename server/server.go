package http

import (
	// stdlib
	"log"
	"net/http"
	// universe
	"github.com/universelabs/universe-server/internal/config"
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
func NewServer(port string, ks *universe.Keystore) *Server {
	srv := &Server {
		// set addr and port
		Addr: port
		// init handler
		Handler: NewHandler(ks)
	}

	// print all routes
	walkFunc := func(method, route string, handler http.Handler, 
		middlewares ...func(http.Handler) http.Handler) error {
			log.Printf("%s -> %s\n", route, method)
			return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error()) // panic if there's an error
	}
}

// Listens and serves the server instance
func (srv *Server) Open() error {
	// open the socket
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	srv.ln = ln

	// start HTTP server
	go func() { http.Serve(srv.ln, srv.Handler) }()
	// *** because http.Serve is called in a go routine, main() must hang
	// the process so that the server doesn't close!

	return nil
}

// Closes the server instance
func (srv *Server) Close() error {
	if srv.ln != nil {
		return srv.ln.Close()
	}
	return nil
}