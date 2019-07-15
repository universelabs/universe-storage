package universe

// Wallet errors
const (
	ErrWalletRequired	= Error("wallet required")
	ErrWalletNotFound	= Error("wallet not found")
	ErrWalletExists		= Error("wallet already exists")
	ErrWalletIDRequired	= Error("wallet ID required")
)

// universe's representation of an error
type Error string

// Implement the error interface
func (e Error) Error() string { return string(e) }
