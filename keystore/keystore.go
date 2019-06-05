package main

import (
	"fmt"
	"log"
	"github.com/asdine/storm"
)

type Keystore struct {
	db *storm.DB
}

type WalletID struct {
	ID int 'storm:"unique, increment"' // uid for stormdb
	Platform string 'storm:"index"' // e.g. ETH, BTC, Blockstack
	Description string 'storm:"index"' // e.g. BTC:personal, ETH:app1, arbitrary descriptor
}

type Wallet struct {
	PublicKey string 
	PrivateKey string 
	Metadata WalletID 'storm:"inline"'
}

func main() {
	var keystore *Keystore
	keystore.db, err := storm.Open("keys.db")
	if err != nil {
		log.Fatal(err)
	}
	defer keystore.db.Close()
	
	// make a few wallet structs for testing

	// TODO: keystore_test.go
}

// CRUD interface

func (keystore *Keystore) AddWallet() err {
	
}

func (keystore *Keystore) RemoveWallet() err {
	
}

func (keystore *Keystore) UpdateWallet() err {
	
}

func (keystore *Keystore) DeleteWallet() err {
	
}