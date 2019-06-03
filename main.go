package main

import (
	"fmt"

	"github.com/asdine/storm"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {

	//Test cause ethereum address wasn't printing
	fmt.Println("Hello giant chicken")

	//Generate Bitcoin address  -- From: https://godoc.org/github.com/btcsuite/btcutil/hdkeychain#example-NewMaster
	// Generate a random seed at the recommended length.
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Show that the generated master node extended key is private.
	fmt.Println("Private Extended Key?:", key.IsPrivate())
	fmt.Println(key)

	//Generate Ethereum Address
	/*
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
			fmt.Print(privateKey)

	*/

	// Storage
	db, err := storm.Open("my.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
