package http

import (
	// universe
	"github.com/universelabs/universe-server/internal/config"
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
	KeystoreHandler KeystoreHandler
}

// Instantiates the chi.Mux and mounts the service handlers for chi's ServeHTTP 
func NewHandler(ks universe.Keystore) *Handler {
	h.Mux = chi.NewRouter()
	h.Mux.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// instantiate and route handlers
	KeystoreHandler = NewKeystoreHandler(ks)
	// don't hardcode the version!
	h.Route("/" /*+ versionString*/, func(r chi.Router) {
		r.Mount("/api/keystore", h.KeystoreHandler.ServeHTTP(w, r))
	})
}
