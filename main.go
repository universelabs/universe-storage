package main

import (
	// universe
	"github.com/unvierselabs/universe-server/universe"
	"github.com/unvierselabs/universe-server/http"
	"github.com/unvierselabs/universe-server/stormdb"
	"github.com/universelabs/universe-server/internal/config"
)

func main() {
	// init config
	consts := config.NewConstants()
	// set up services

	// wire everything together

	// launch server
}

// func main() {
// 	// HANDLE CFG
// 	var err error
// 	var cfg config.Config
// 	if cfg, err = config.New(); err != nil {
// 		log.Panicf("Configuration error: %v\n", err)
// 	}

// 	router := Routes(cfg)
	
// 	// PART OF SERVER.NEWCLIENT() OR SMTHING
// 	// print all routes
// 	walkFunc := func(method, route string, handler http.Handler, 
// 		middlewares ...func(http.Handler) http.Handler) error {
// 			log.Printf("%s -> %s\n", route, method)
// 			return nil
// 	}
// 	if err := chi.Walk(router, walkFunc); err != nil {
// 		log.Panicf("Logging error: %s\n", err.Error()) // panic if there's an error
// 	}
	
// 	log.Fatal(http.ListenAndServe(":8080", router)) // **port should be from env not hardcoded
// }