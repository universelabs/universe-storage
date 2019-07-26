package http

import (
	// stdlib
	// "net/http"
	// "log"
	// universe
	"github.com/universelabs/universe-server/universe"
	// deps
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Handler is a collection of all the service handlers
type Handler struct {
	// router
	*chi.Mux
	// service handlers
	KeystoreHandler *KeystoreHandler
}

// Instantiates the chi.Mux and mounts the service handlers for chi's ServeHTTP 
func NewHandler(ks universe.Keystore) *Handler {
	h := &Handler{}
	h.Mux = chi.NewRouter()
	h.Mux.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// instantiate and route handlers
	h.KeystoreHandler = NewKeystoreHandler(ks)
	// don't hardcode the version!
	h.Route("/v" /*+ versionString*/, func(r chi.Router) {
		r.Mount("/api/keystore", h.KeystoreHandler)
	})

	// 	// print all routes
	// walkFunc := func(method, route string, handler http.Handler, 
	// 	middlewares ...func(http.Handler) http.Handler) error {
	// 		log.Printf("[Handler] %s -> %s\n", route, method)
	// 		return nil
	// }
	// if err := chi.Walk(h.Mux, walkFunc); err != nil {
	// 	log.Panicf("[Handler] Logging error: %s\n", err.Error()) // panic if there's an error
	// }


	return h
}
