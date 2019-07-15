package internal

import (
	"encoding/json"
)

// Unmarshalling wallets from JSON
func (w *universe.Wallet) UnmarshalJSON(data []byte) error {
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
		key := universe.BlockstackID{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err 
		}
		w.Data = key
	case "Ethereum":
		key := universe.ETHKey{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err
		}
		w.Data = key
	case "Bitcoin":
		key := universe.BTCKey{}
		if err = json.Unmarshal(*objMap["Data"], &key); err != nil {
			return err
		}
		w.Data = key
	}
	
	return nil
}