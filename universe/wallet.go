package universe

// A wallet is represented as metadata and a platform implementation
type Wallet struct {
	ID int `storm:"id,increment"`
	Platform string `storm:"index"` // e.g. ETH, BTC, Blockstack
	Description string `storm:"index"` // e.g. BTC:personal, ETH:app1, arbitrary descriptor

	Data Key `storm:"inline"` 
}

// Different platforms must implement this (at the moment an "empty" interface)
type Key interface {
	isKey()
}

// Represents a Blockstack ID
type BlockstackID struct {
	UID string
	Email string
	Password string
	Passphrase string 
}

func (b BlockstackID) isKey() {}

// Represents an Ethereum wallet
type ETHKey struct {
	PublicKey string
	PrivateKey string
}

func (b ETHKey) isKey() {}

// Represents a Bitcoin wallet
type BTCKey struct {
	PublicKey string
	PrivateKey string 
}

func (b BTCKey) isKey() {}
