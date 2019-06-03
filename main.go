package main

import (
	"fmt"
	"log"

	"github.com/asdine/storm"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	//Ethereum Private Key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(privateKey)

	// Storage
	db, err := storm.Open("my.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
