package storage

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetAllWallets)
	// router.Get("/platform/{platformID}", GetPlatform)
	// router.Get("/key/{keyID}", GetKey)
	return router
}

func GetAllWallets(w http.ResponseWriter, r *http.Request) {
	// testing the route
	// wallets := []Wallet{
	// 	{ID: 0,
	// 	Platform: "Blockstack",
	// 	Description: "test1",
	// 	Data: BlockstackID{
	// 		UID: "123",
	// 		Email: "a@b.com",
	// 		Password: "12345",
	// 		Passphrase: "hello-bye",},},
	// 	{ID: 1,
	// 	Platform: "Ethereum",
	// 	Description: "test1",
	// 	Data: ETHKey{
	// 		PublicKey: "wiq73yrh79yr9rf93hfyca",
	// 		PrivateKey: "fgbosfgnuonoufnduonf3f3o",},}}
	// render.JSON(w, r, wallets)
}