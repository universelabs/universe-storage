package http_test

import (
	"github.com/universelabs/universe-server/http"
)

// test wrapper for http.Handler
type Handler struct {
	*http.Handler
	// test wrapper for KeystoreHandler
	KeystoreHandler *KeystoreHandler
}

// returns a new instance of Handler
func NewHandler() *Handler {
	h := &Handler {
		Handler: &http.Handler{},
		KeystoreHandler: NewKeystoreHandler{},
	}
	h.Handler.KeystoreHandler = h.KeystoreHandler.KeystoreHandler
	return h
}