package rk


// Account holds address, balance and public key information
type Account struct {
	address string
	balance float32
	txs int
}

// map of address to Account
type Accounts map[string] Account

// Create New Account
func NewAccount() *Account {
	account := new(Account)
	account.address = GenerateAddressString()
	return account
}
