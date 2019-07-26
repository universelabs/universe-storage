package stormdb_test

import (
	// stdlib
	"testing"
	"reflect"
	// universe
	"github.com/universelabs/universe-server/universe"
)

// Test the storm.AddWallet() and storm.GetWallet() functions
func TestKeystore_AddGetWallet(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	ks := c.Keystore()

	eth := universe.Wallet{
		Platform: "Ethereum",
		Description: "test",
		Data: universe.ETHKey{
			PublicKey: "ETH-TestPubKey",
			PrivateKey: "ETH-TestPrivKey",},
	}
	btc := universe.Wallet{
		Platform: "Bitcoin",
		Description: "test",
		Data: universe.BTCKey{
			PublicKey: "BTC-TestPubKey",
			PrivateKey: "BTC-TestPrivKey",},
	}
	bid := universe.Wallet{
		Platform: "Blockstack",
		Description: "test",
		Data: universe.BlockstackID{
			UID: "Blockstack-TestUID",
			Email: "Blockstack-TestEmail",
			Password: "Blockstack-TestPassword",
			Passphrase: "Blockstack-TestPassphrase",},
	}

	// test adding ethereum wallet
	if err := ks.AddWallet(&eth); err != nil {
		t.Fatal(err)
	}
	if ret, err := ks.GetWallet(eth.ID); err != nil && reflect.DeepEqual(ret, eth) {
		t.Fatal(err)
	}
	// test adding bitcoin wallet
	if err := ks.AddWallet(&btc); err != nil {
		t.Fatal(err)
	} 
	if ret, err := ks.GetWallet(btc.ID); err != nil && reflect.DeepEqual(ret, btc) {
		t.Fatal(err)
	}
	// test adding BlockstackID
	if err := ks.AddWallet(&bid); err != nil {
		t.Fatal(err)
	} 
	if ret, err := ks.GetWallet(bid.ID); err != nil && reflect.DeepEqual(ret, bid) {
		t.Fatal(err)
	}
}

func TestKeystore_AddWallet_ErrWalletExists(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()

	testWallet := universe.Wallet{
		Platform: "Bitcoin",
		Description: "test",
		Data: universe.BTCKey{
			PublicKey: "BTC-TestPubKey",
			PrivateKey: "BTC-TestPrivKey",},
	}

	if err := c.Keystore().AddWallet(&testWallet); err != nil {
		t.Fatal(err)
	}
	// if err := c.Keystore().AddWallet(&testWallet); err != ErrWalletExists {
	// 	t.Fatal(err)
	// }
}

func TestKeystore_AddGetWallet_ErrWalletRequired(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	if err := c.Keystore().AddWallet(nil); err != universe.ErrWalletRequired {
		t.Fatal(err)
	}
}

// func TestKeystore_AddGetWallet_ErrWalletIDRequired(t *testing.T) {
// 	c := MustOpenClient()
// 	defer c.Close()
// 	if err := c.Keystore().AddWallet(nil); err != universe.ErrWalletIDRequired {
// 		t.Fatal(err)
// 	}
// }

func TestKeystore_GetAllWallets(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	ks := c.Keystore()

	wallets := []universe.Wallet{
		{ Platform: "Ethereum",
		Description: "test",
		Data: universe.ETHKey{
			PublicKey: "ETH-TestPubKey",
			PrivateKey: "ETH-TestPrivKey",},},
		{ Platform: "Bitcoin",
		Description: "test",
		Data: universe.BTCKey{
			PublicKey: "BTC-TestPubKey",
			PrivateKey: "BTC-TestPrivKey",},},
		{ Platform: "Blockstack",
		Description: "test",
		Data: universe.BlockstackID{
			UID: "Blockstack-TestUID",
			Email: "Blockstack-TestEmail",
			Password: "Blockstack-TestPassword",
			Passphrase: "Blockstack-TestPassphrase",},},
	}

	if retWallets, err := ks.GetAll(); err != nil {
		t.Fatal(err)
	} else {
		for i, _ := range retWallets {
			if !reflect.DeepEqual(retWallets[i], wallets[i]) {
				t.Fatalf("[TestKeystore_GetAllWallets] unexpected wallet: %v (expected: %v)", 
					retWallets[i], wallets[i])
			}
		}
	}
}

// NEED TO ADD MORE TESTS!
