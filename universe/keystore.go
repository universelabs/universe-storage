package universe

type Client interface {
	Keystore() Keystore
}
	
type Keystore interface {
	AddWallet(new_wallet *Wallet) error
	GetWallet(id int) (*Wallet, error)
	GetPlatform(platform string) ([]Wallet, error)
	GetAll() ([]Wallet, error)
	DeleteWallet(id int) error
}