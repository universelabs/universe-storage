package storage

import (
	"encoding/json"
)

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