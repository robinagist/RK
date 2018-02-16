package rk

import (
	"time"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"encoding/hex"
	"math/rand"
)

// Account holds address, balance and public key information
type Account struct {
	address string
	balance float32
	txs int
}

// map of address to Account
type Accounts map[string] *Account

// Create New Account
func (acs *Accounts) NewAccount() *Account {
	account := new(Account)
	account.address = GenerateAddressString()
	return account
}

func GenerateAddressString() string {
	t := string(time.Now().Nanosecond())
	n := sha3.New256()
	n.Write([]byte(t + string(rand.Intn(1000000))))
	return ("099" + hex.EncodeToString(n.Sum(nil)))[0:31]
}