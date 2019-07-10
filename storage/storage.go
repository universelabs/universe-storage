package storage

import (
	"github.com/asdine/storm"
)
	
type Keystore struct {
	db *storm.DB
}

// keystore methods

func (ks *Keystore) Init(pathname string, newInstance bool) (err error) {
	// check against newInstance
	ks.db, err = storm.Open(pathname)
	if err != nil {
		return err
	}
	err = ks.db.Init(&Wallet{})
	return err
}

func (ks *Keystore) Close() error {
	err := ks.db.Close()
	return err
}

// CRUD interface

func (ks *Keystore) AddWallet(new_wallet *Wallet) error {
	err := ks.db.Save(new_wallet)
	return err
}

func (ks *Keystore) GetWallet(id int) (Wallet, error) {
	wallet := Wallet{}
	err := ks.db.One("ID", id, &wallet)
	return wallet, err
}

func (ks *Keystore) GetPlatform(platform string) ([]Wallet, error) {
	wallets := []Wallet{}
	err := ks.db.Find("Platform", platform, &wallets)
	return wallets, err
}

func (ks *Keystore) GetAll() ([]Wallet, error) {
	wallets := []Wallet{}
	err := ks.db.All(&wallets)
	return wallets, err
}

func (ks Keystore) DeleteWallet(id int) error {
	get, geterr := ks.GetWallet(id)
	if geterr != nil {
		return geterr
	} else {
		return ks.db.DeleteStruct(&get)
	}
}

// func main() {
// 	ks := Keystore{}
// 	err := ks.Init("db.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ks.Close()
// 	// make a few wallet structs for testing
// 	eth1 := Wallet{
// 		Platform: "Ethereum",
// 		Description: "test1",
// 		Data: ETHKey{
// 			PublicKey: "wiq73yrh79yr9rf93hfyca",
// 			PrivateKey: "fgbosfgnuonoufnduonf3f3o",},
// 	}
// 	bid := Wallet{
// 		Platform: "Blockstack",
// 		Description: "test1",
// 		Data: BlockstackID{
// 			UID: "123",
// 			Email: "a@b.com",
// 			Password: "12345",
// 			Passphrase: "hello-bye",},
// 	}
	
// 	saveerr := ks.AddWallet(&bid)
// 	if saveerr != nil {
// 		fmt.Println(saveerr)
// 	} else {
// 		fmt.Printf("saved bid:%v\n", bid)
// 	}
// 	saveerr = ks.AddWallet(&eth1)
// 	if saveerr != nil {
// 		fmt.Println(saveerr)
// 	} else {
// 		fmt.Printf("saved bid:%v\n", eth1)
// 	}
// 	ret, reterr := ks.GetAll()
// 	if reterr != nil {
// 		fmt.Printf("err: %q\n", reterr)	
// 	} else {
// 		fmt.Println(ret)
// 	}
// 	// TODO: keystore_test.go
// }