package http

import (
	// stdlib
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	// universe
	"github.com/unvierselabs/universe-server/universe"
	// deps
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Keystore Handler is the HTTP API for key storage
type KeystoreHandler struct { 
	// router
	*chi.Mux
	// services
	Keystore universe.Keystore
	// utilities
	Logger *log.Logger
}

// NewKeystoreHandler returns a new instance of KeystoreHandler
func NewKeystoreHandler(ks universe.Keystore) *KeystoreHandler {
	h := &KeystoreHandler{
		Mux: chi.NewRouter(),
		Keystore: ks, // FIX 
		Logger: log.New(os.Stderr, "[Keystore] ", log.LstdFlags)
	}
	h.Mux.Post("/addwallet", AddWallet())
	h.Mux.Get("/wallet/{walletID}", GetWallet())
	h.Mux.Get("/platform/{platformID}", GetPlatform())
	h.Mux.Get("/", GetAll())

	// print all routes
	walkFunc := func(method, route string, handler http.Handler, 
		middlewares ...func(http.Handler) http.Handler) error {
			log.Printf("[KeystoreHandler] %s -> %s\n", route, method)
			return nil
	}
	if err := chi.Walk(h.Mux, walkFunc); err != nil {
		log.Panicf("[KeystoreHandler] Logging error: %s\n", err.Error()) // panic if there's an error
	}

	return h
}

// Adds a wallet to the keystore from the HTTP request and returns a confirmation
func (ksh *KeystoreHandler) AddWallet(w http.ResponseWriter, r *http.Request) {
		// unmarshal from json
		wallet := Wallet{}
		// log.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&wallet)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("added wallet: %v", wallet)
		}
		// add to db
		err = ksh.Keystore.AddWallet(&wallet)
		// return err if failed, else confirmation
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, wallet)
		}
	}
}

// REST API wrapper for universe.Keystore.GetWallet()
func (ksh *KeystoreHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
		walletID, urlerr := strconv.Atoi(chi.URLParam(r, "walletID"))
		if urlerr != nil {
			log.Println(urlerr)
			render.JSON(w, r, urlerr)
			return
		}
		ret, err := ksh.Keystore.GetWallet(walletID)
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}

// REST API wrapper for universe.Keystore.GetPlatform()
func (ksh *KeystoreHandler) GetPlatform(w http.ResponseWriter, r *http.Request) {
		platformID := chi.URLParam(r, "platformID")
		ret, err := ksh.Keystore.GetPlatform(platformID)
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}

// REST API wrapper for universe.Keystore.GetAll()
func (ksh *KeystoreHandler) GetAll(w http.ResponseWriter, r *http.Request) {
		ret, err := ksh.Keystore.GetAll()
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}
