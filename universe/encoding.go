package universe

import (
	// stdlib
	"encoding/json"
	"strconv"
)

func (w *Wallet) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["id"] = strconv.Itoa(w.ID)
	m["platform"] = w.Platform
	m["description"] = w.Description
	if key, err := json.Marshal(&w.Data); err != nil {
		return []byte{}, err
	} else {
		m["data"] = string(key)
	}
	return json.Marshal(m)
}

// Unmarshalling wallets from JSON
func (w *Wallet) UnmarshalJSON(data []byte) error {
	// deconstruct the serialized json into a map
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objMap)
	if err != nil {
		return err
	}

	// unmarshal the metadata
	if idptr, ok := objMap["id"]; ok {
		if err = json.Unmarshal(*idptr, &w.ID); err != nil {
			return err 
		}
	}
	
	if err = json.Unmarshal(*objMap["platform"], &w.Platform); err != nil {
		return err
	}

	if err = json.Unmarshal(*objMap["description"], &w.Description); err != nil {
		return err
	}
	
	// instantiate the correct Key struct
	switch w.Platform {
	case "Blockstack":
		key := BlockstackID{}
		if err = json.Unmarshal(*objMap["data"], &key); err != nil {
			return err 
		}
		w.Data = key
	case "Ethereum":
		key := ETHKey{}
		if err = json.Unmarshal(*objMap["data"], &key); err != nil {
			return err
		}
		w.Data = key
	case "Bitcoin":
		key := BTCKey{}
		if err = json.Unmarshal(*objMap["data"], &key); err != nil {
			return err
		}
		w.Data = key
	}

	return nil
}