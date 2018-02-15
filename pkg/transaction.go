package rk

import (
	"time"
	"errors"
)

//// Transactions
type Transaction struct {
	TxId int              `json:"txId"`
	Timestamp time.Time   `json:"timestamp"`
	Ttl int               `json:"ttl"`
	Sender string         `json:"sender"`
	Recipient string      `json:"recipient"`
	TxType int            `json:"txType"`
	Data string           `json:"data"`
	Amount float32        `json:"amount"`
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
	data string) (*Transaction, error) {

	// make sure transaction is valid
    if !prevalidate(sender, recipient, amount) {
    	return nil, errors.New("amount exceeds sender's account balance")
	}

	// create the transaction
	transaction := Transaction {
		TxId: sender.txs,
		Sender: sender.address,
		Recipient: recipient.address,
		TxType: txType,
		Amount: amount,
		Data: data,
	}

	if amount > 0 {
		transaction.transfer(sender, recipient, amount)
	}
	return &transaction, nil
}


func prevalidate(sender *Account, recipient *Account, amount float32) bool {
	if sender.balance < amount {
		return false
	}
	return true
}