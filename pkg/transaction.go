package rk

import (
	"time"
	"errors"
)

//// Transactions
type Transaction struct {
	txId int
	timestamp time.Time
	ttl int
	sender string
	recipient string
	txType int
	data []byte
	amount float32
}

// transfer from one account to another
func (t *Transaction) transfer (sender *Account, recipient *Account, amount float32) {
	sender.balance -= amount
	recipient.balance += amount
}

// Creates a new transaction
func NewTransaction(
	sender *Account,
	recipient *Account,
	txType int,
	amount float32,
	data []byte) (*Transaction, error) {

	// make sure transaction is valid
    if !prevalidate(sender, recipient, amount) {
    	return nil, errors.New("amount exceeds sender's account balance")
	}

	// create the transaction
	transaction := Transaction {
		txId: sender.txs,
		sender: sender.address,
		recipient: recipient.address,
		txType: txType,
		amount: amount,
		data: data,
	}

	if amount > 0 {
		transaction.transfer(sender, recipient, amount)
	}
	return &transaction, nil
}


// Transaction Pool
type TransactionPool struct {
	tPool []Transaction
}

func (tp *TransactionPool) filter(criteria string) []Transaction {
	return nil
}


func prevalidate(sender *Account, recipient *Account, amount float32) bool {
	if sender.balance < amount {
		return false
	}
	return true
}