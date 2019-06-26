package storage

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	keystore *Keystore
)

func Routes(ks *Keystore) *chi.Mux {
	router := chi.NewRouter()
	keystore = ks
	router.Post("/addwallet", AddWallet)
	router.Get("/wallet/{walletID}", GetWallet)
	router.Get("/platform/{platformID}", GetPlatform)
	router.Get("/", GetAllWallets)
	return router
}

func AddWallet(w http.ResponseWriter, r *http.Request) {
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
	err = keystore.AddWallet(&wallet)
	// return err if failed, else confirmation
	if err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, wallet)
	}
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	walletID, urlerr := strconv.Atoi(chi.URLParam(r, "walletID"))
	if urlerr != nil {
		log.Println(urlerr)
		render.JSON(w, r, urlerr)
		return
	}
	ret, err := keystore.GetWallet(walletID)
	if err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, ret)
	}
}

func GetPlatform(w http.ResponseWriter, r *http.Request) {
	platformID := chi.URLParam(r, "platformID")
	ret, err := keystore.GetPlatform(platformID)
	if err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, ret)
	}
}

func GetAllWallets(w http.ResponseWriter, r *http.Request) {
	ret, err := keystore.GetAll()
	if err != nil {
		render.JSON(w, r, err)
	} else {
		render.JSON(w, r, ret)
	}
}

