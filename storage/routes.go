package storage

import (
	"log"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/universelabs/universe-server/internal/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes(cfg *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/addwallet", AddWallet(cfg))
	router.Get("/wallet/{walletID}", GetWallet(cfg))
	router.Get("/platform/{platformID}", GetPlatform(cfg))
	router.Get("/", GetAllWallets(cfg))
	return router
}

func AddWallet(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		err = cfg.KS.AddWallet(&wallet)
		// return err if failed, else confirmation
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, wallet)
		}
	}
}

func GetWallet(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, urlerr := strconv.Atoi(chi.URLParam(r, "walletID"))
		if urlerr != nil {
			log.Println(urlerr)
			render.JSON(w, r, urlerr)
			return
		}
		ret, err := cfg.KS.GetWallet(walletID)
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}

func GetPlatform(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		platformID := chi.URLParam(r, "platformID")
		ret, err := cfg.KS.GetPlatform(platformID)
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}

func GetAllWallets(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ret, err := cfg.KS.GetAll()
		if err != nil {
			render.JSON(w, r, err)
		} else {
			render.JSON(w, r, ret)
		}
	}
}
