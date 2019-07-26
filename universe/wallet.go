package universe

// A wallet is represented as metadata and a platform implementation
type Wallet struct {
	ID int `storm:"id,increment"`
	Platform string `"storm:"index"` // e.g. ETH, BTC, Blockstack
	Description string `"storm:"index"` // e.g. BTC:personal, ETH:app1, arbitrary descriptor

	Data Key `storm:"inline"` 
}

// Different platforms must implement this (at the moment an "empty" interface)
type Key interface {
	isKey()
}

// Represents a Blockstack ID
type BlockstackID struct {
	UID string `json:"uid"`
	Email string `json:"email"`
	Password string `json:"password"`
	Passphrase string `json:"Passphrase"` 
}

func (b BlockstackID) isKey() {}

// Represents an Ethereum wallet
type ETHKey struct {
	PublicKey string `json:"publickey"`
	PrivateKey string `json:"privatekey"`
}

func (b ETHKey) isKey() {}

// Represents a Bitcoin wallet
type BTCKey struct {
	PublicKey string `json:"publickey"`
	PrivateKey string `json:"privatekey"` 
}

func (b BTCKey) isKey() {}
