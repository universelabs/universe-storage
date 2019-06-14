package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	
	"keystore"
)

var (
	storage keystore.Keystore
)

// "/" will only return metadata
// "/{platform}" will return data for specific platform
// "/all" 

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/0.0.7", func(r chi.Router) {
		r.Mount("/api/keystore", keystore.Routes()
	})
	return router
}

func handleRequests() {
	http.HandleFunc("/", )
}

func main() {
	storage = keystore.storage.Keystore{}
	storage.Init("keys.db")

	router := Routes()
	
	// print all routes
	walkFunc := func(method, route string, handler http.Handler, 
		middlewares ...func(http.Handler) http.Handler) error {
			log.Printf("%s -> %s\n", route, method)
			return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error()) // panic if there's an error
	}

	log.Fatal(http.ListenAndServer(":8080", router)) // **port should be from env not hardcoded
}