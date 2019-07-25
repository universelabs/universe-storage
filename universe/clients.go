package universe

// A client with which to interact with the Keystore service
type KeystoreClient interface {
	Keystore() Keystore
}
	
// Keystore service interface
type Keystore interface {
	AddWallet(new_wallet *Wallet) error
	GetWallet(id int) (Wallet, error)
	GetPlatform(platform string) ([]Wallet, error)
	GetAll() ([]Wallet, error)
	DeleteWallet(id int) error
}