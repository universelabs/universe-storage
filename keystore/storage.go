package main

import (
	"fmt"
	"log"
	// "strconv"
	"encoding/json"

	"github.com/asdine/storm"

	// "github.com/universelabs/universe-server/keystore"
)
	
type Keystore struct {
	db *storm.DB
}

// CRUD interface

func (keystore *Keystore) AddWallet(new_wallet *Wallet) error {
	err := keystore.db.Save(new_wallet)
	return err
}

func (keystore *Keystore) GetWallet(id int) (Wallet, error) {
	wallet := Wallet{}
	err := keystore.db.One("ID", id, &wallet)
	return wallet, err
}

func (keystore *Keystore) GetPlatform(platform string) ([]Wallet, error) {
	wallets := []Wallet{}
	err := keystore.db.Find("Platform", platform, &wallets)
	return wallets, err
}

func (keystore *Keystore) GetAll() ([]Wallet, error) {
	wallets := []Wallet{}
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

func main() {
	keystore := Keystore{}
	err := keystore.Init("db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer keystore.Close()
	// make a few wallet structs for testing
	// eth1 := Wallet{
	// 	Platform: "Ethereum",
	// 	Description: "test1",
	// 	Data: ETHKey{
	// 		PublicKey: "wiq73yrh79yr9rf93hfyca",
	// 		PrivateKey: "fgbosfgnuonoufnduonf3f3o",},
	// }
	bid := Wallet{
		ID: 1,
		Platform: "Blockstack",
		Description: "test1",
		Data: BlockstackID{
			UID: "123",
			Email: "a@b.com",
			Password: "12345",
			Passphrase: "hello-bye",},
	}
	
	saveerr := keystore.AddWallet(&bid)
	if saveerr != nil {
		fmt.Println(saveerr)
	} else {
		fmt.Printf("saved bid:%v\n", bid)
	}
	ret, reterr := keystore.GetWallet(1)
	if reterr != nil {
		fmt.Printf("err: %q\n", reterr)	
	} else {
		fmt.Println(ret)
	}
	// TODO: keystore_test.go
}

type Wallet struct {
	ID int `storm:"id,increment"` // uid for stormdb
	Platform string `storm:"index"` // e.g. ETH, BTC, Blockstack
	Description string `storm:"index"` // e.g. BTC:personal, ETH:app1, arbitrary descriptor

	Data Key `storm:"inline"` 
}

type Key interface {
	IsKey()
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

// keystore methods

func (keystore *Keystore) Init(pathname string) (err error) {
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

// Unmarshalling wallets 

func (w *Wallet) UnmarshalJSON(data []byte) error {
	// deconstruct the serialized json into a map
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objMap)
	if err != nil {
		return err
	}
	
	// unmarshal the metadata 
	if err = json.Unmarshal(*objMap["ID"], &w.ID); err != nil {
		fmt.Println(err)
		return err
	}
	
	if err = json.Unmarshal(*objMap["Platform"], &w.Platform); err != nil {
		return err
	}

	if err = json.Unmarshal(*objMap["Description"], &w.Description); err != nil {
		return err
	}
	
	// instantiate the correct Key struct
	switch w.Platform {
	case "Blockstack":
		key := BlockstackID{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err 
		}	
		fmt.Println(key)
		w.Data = key
	case "Ethereum":
		key := ETHKey{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err
		}
		w.Data = key
	case "Bitcoin":
		key := BTCKey{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err
		}
		w.Data = key
	}
	
	return nil
}