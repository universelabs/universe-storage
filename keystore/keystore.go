package main

import (
	"fmt"
	"log"
	// "os"
	// "bufio"
	"github.com/asdine/storm"
)

type Keystore struct {
	db *storm.DB
}

type Key interface {
	IsKey()
}

type Wallet struct {
	ID int `storm:"id,increment"` // uid for stormdb
	Platform string `storm:"index"` // e.g. ETH, BTC, Blockstack
	Description string `storm:"index"` // e.g. BTC:personal, ETH:app1, arbitrary descriptor

	Data Key `storm:"inline"` 
}

type BlockstackID struct {
	UID string
	Email string
	Password string
	Passphrase string 
}

func (b BlockstackID) IsKey() {}

type ETHKey struct {
	PublicKey string
	PrivateKey string
}

func (b ETHKey) IsKey() {}

type BTCKey struct {
	PublicKey string
	PrivateKey string 
}

func (b BTCKey) IsKey() {}

func main() {
	keystore := Keystore{}
	err := keystore.Init("db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer keystore.Close()
	// make a few wallet structs for testing
	eth1 := Wallet{
		Platform: "eth",
		Description: "test1",
		Data: ETHKey{
			PublicKey: "wiq73yrh79yr9rf93hfyca",
			PrivateKey: "fgbosfgnuonoufnduonf3f3o",},
	}
	fmt.Println(eth1)
	saveerr := keystore.AddWallet(&eth1)
	if saveerr != nil {
		fmt.Println(saveerr)
	} else {
		fmt.Printf("saved eth1:%v\n", eth1)
	}
	ret, reterr := keystore.GetPlatform("eth")
	if reterr != nil {
		fmt.Println(reterr)
	} else {
		fmt.Println(ret)
	}
	// TODO: keystore_test.go
}

// keystore methods

func (keystore *Keystore) Init(pathname string) error {
	var err error
	keystore.db, err = storm.Open(pathname)
	if err != nil {
		return err
	}
	err = keystore.db.Init(&Wallet{})
	return err
}

func (keystore *Keystore) Close() error {
	err := keystore.db.Close()
	return err
}

// CRUD interface

func (keystore *Keystore) AddWallet(new_wallet *Wallet) error {
	err := keystore.db.Save(new_wallet)
	return err
}

func (keystore *Keystore) GetWallet(id int) (Wallet, error) {
	var wallet Wallet
	err := keystore.db.One("ID", id, &wallet)
	return wallet, err
}

func (keystore *Keystore) GetPlatform(platform string) ([]Wallet, error) {
	var wallets []Wallet
	err := keystore.db.Find("Platform", platform, &wallets)
	return wallets, err
}

func (keystore *Keystore) GetAll() ([]Wallet, error) {
	var wallets []Wallet
	err := keystore.db.All(&wallets)
	return wallets, err
}

// func (keystore Keystore) UpdateWallet(id int, new_wallet *Wallet) error {
// 	err := keystore.db.UpdateField()
// 	return err
// }

func (keystore Keystore) DeleteWallet(id int) error {
	get, geterr := keystore.GetWallet(id)
	if geterr != nil {
		return geterr
	} else {
		return keystore.db.DeleteStruct(&get)
	}
}