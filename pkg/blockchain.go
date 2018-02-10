package rk

import (
	"time"
	"encoding/json"
	"crypto/sha256"
)

type Transaction struct {
	sender string
	recipient string
	txType int
	data []byte
	amount float32
}

type Block struct {
	index uint64
	timestamp time.Time
	transactions []Transaction
	proof uint
	previousHash string
}

type BlockChain struct {
    txQueue []Transaction
    chain []Block
}


// creates a new block on BlockChain
func (bc *BlockChain) NewBlock (proof uint, previousHash string) *Block {

	indexCount := uint64(len(bc.chain))
	timestamp := time.Time{}
	block := Block {
		index: indexCount,
		proof: proof,
		transactions: bc.txQueue,
		timestamp: timestamp,
		previousHash: previousHash,
		}

	bc.txQueue = nil
	bc.chain = append(bc.chain, block)
	return &block
}


// creates a new transaction
func (bc *BlockChain) NewTransaction (sender string, recipient string, txType int, amount float32, data []byte) int {

	transaction := Transaction {
		sender: sender,
		recipient: recipient,
		txType: txType,
		amount: amount,
		data: data,
	}

	bc.txQueue = append(bc.txQueue, transaction)

	return len(bc.chain)
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
   hash := h.Sum(pfx)
   lh := len(hash)

   // get the last three characters
   matching := string(hash[lh-3:lh-1])

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
	hash := h.Sum(blockJson)
	return hash
}










