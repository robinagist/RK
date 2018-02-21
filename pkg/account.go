package rk

import (
	"time"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"encoding/hex"
	"math/rand"
)

// Account holds address, balance and public key information
type Account struct {
	Address string     `json:"address"`
	Balance float32    `json:"balance"`
	Txs int            `json:"txs"`
}

// map of address to Account
type Accounts map[string] *Account

// Create New Account
func (acs *Accounts) NewAccount(address string) *Account {
	account := new(Account)
	if address == "" {
		account.Address = GenerateAddressString()
	} else { account.Address = address}
	return account
}

func GenerateAddressString() string {
	t := string(time.Now().Nanosecond())
	n := sha3.New256()
	n.Write([]byte(t + string(rand.Intn(1000000))))
	return ("099" + hex.EncodeToString(n.Sum(nil)))[0:31]
}