package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	
	"github.com/universelabs/universe-server/storage"
)

var (
	keystore storage.Keystore
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

	router.Route("/0.0.10", func(r chi.Router) {
		r.Mount("/api/keystore", storage.Routes(&keystore))
	})
	return router
}

func main() {
	keystore = storage.Keystore{}
	keystore.Init("keys.db")

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

	// keystore.AddWallet(&storage.Wallet{
	// 	Platform: "Ethereum",
	// 	Description: "test1",
	// 	Data: storage.ETHKey{
	// 		PublicKey: "wiq73yrh79yr9rf93hfyca",
	// 		PrivateKey: "fgbosfgnuonoufnduonf3f3o",},
	// })

	// ret, err := keystore.GetWallet(1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(ret)
	// }

	log.Fatal(http.ListenAndServe(":8080", router)) // **port should be from env not hardcoded
}