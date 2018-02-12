package rk

import (
	"time"
	"encoding/json"
	"crypto/sha256"
	"math/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"errors"
)

type Account struct {
	address string
	balance float32
	txs int
}

type Transaction struct {
	txId int
	sender string
	recipient string
	txType int
	data []byte
	amount float32
}

type Block struct {
	index int
	timestamp time.Time
	transactions []Transaction
	proof uint
	previousHash string
}

type BlockChain struct {
    txQueue []Transaction
    chain []Block
}

// map of address to Account
type Accounts map[string] Account

// transaction queue
var TXQueue []Transaction

func init() {
	rand.Seed(time.Now().UnixNano())
}

// creates a new block on BlockChain
func (bc *BlockChain) NewBlock (proof uint, previousHash string) *Block {

	indexCount := len(bc.chain)
	timestamp := time.Time{}

	block := Block {
		index: indexCount,
		proof: proof,
		transactions: TXQueue,
		timestamp: timestamp,
		previousHash: previousHash,
		}

	TXQueue = nil
	bc.chain = append(bc.chain, block)
	return &block
}


// creates a new transaction
func NewTransaction(
	sender *Account,
	recipient *Account,
	txType int,
	amount float32,
	data []byte) (*Transaction, error) {

	// make sure transaction is valid
	// TODO - move this out to its own method
	if amount > sender.balance {
		return nil, errors.New("not enough in sender balance to complete transaction")
	}


	sender.balance -= amount
	sender.txs += 1

		transaction := Transaction {
			txId: sender.txs,
			sender: sender.address,
			recipient: recipient.address,
			txType: txType,
			amount: amount,
			data: data,
		}

	TXQueue = append(TXQueue, transaction)

    return &transaction, nil
}


// proof of work algorithm
func (bc *BlockChain) ProofOfWork (lastProof int) int {

	proof := 0
	for {
		if ValidProof(lastProof, proof) {
			return proof
		}
	proof += 1
	}
}

// validate proof
func ValidProof(lastProof int, proof int) bool {

   pfx := []byte(string(lastProof) + string(proof))
   h := sha256.New()
   hsh := h.Sum(pfx)
   lh := len(hsh)

   // get the last three characters
   matching := string(hsh[lh-3:lh-1])

   // arbitrary matching parameter
   return matching == "000"
}

// hash block
func Hash(block *Block) []byte {

	blockJson, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}

	// get a hex digest
	h := sha256.New()
	return h.Sum(blockJson)
}


// Create Account factory
func NewAccount() *Account {
	account := new(Account)
	account.address = GenerateAddressString()
    return account
}

func GenerateAddressString() string {
	t := string(time.Now().Nanosecond())
	n := sha3.New256()
	n.Write([]byte(t + string(rand.Intn(1000000))))
	return "099" + hex.EncodeToString(n.Sum(nil))
}








