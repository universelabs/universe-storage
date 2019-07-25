package stormdb

import (
	// universe
	"github.com/universelabs/universe-server/universe"
)

// represents the service to service storing wallets
type Keystore struct {
	client *Client
}

// Keystore interface implementation

func (ks *Keystore) AddWallet(new_wallet *universe.Wallet) error {
	// require object and id
	if new_wallet == nil {
		return universe.ErrWalletRequired
	}
	// else if new_wallet.ID == "" {
	// 	return universe.ErrWalletIDRequired
	// }

	err := ks.client.db.Save(new_wallet)
	return err
}

func (ks *Keystore) GetWallet(id int) (universe.Wallet, error) {
	wallet := universe.Wallet{}
	err := ks.client.db.One("ID", id, &wallet)
	if err != nil {
		return nil, err
	}
	return wallet, err
}

func (ks *Keystore) GetPlatform(platform string) ([]universe.Wallet, error) {
	wallets := []universe.Wallet{}
	err := ks.client.db.Find("Platform", platform, &wallets)
	return wallets, err
}

func (ks *Keystore) GetAll() ([]universe.Wallet, error) {
	wallets := []universe.Wallet{}
	err := ks.client.db.All(&wallets)
	return wallets, err
}

func (ks *Keystore) DeleteWallet(id int) error {
	get, geterr := ks.GetWallet(id)
	if geterr != nil {
		return geterr
	} else {
		return ks.client.db.DeleteStruct(&get)
	}
}
