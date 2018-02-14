package rk

import (
	"time"
	"encoding/json"
	"crypto/sha256"
	"math/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"encoding/binary"
	"rk/internal"
)


type Block struct {
	index         int
	nonce         uint64
	timestamp     time.Time
	transactions  []Transaction
	hashPrevBlock string
}

type BlockChain struct {
    txQueue []Transaction
    chain []Block
    target uint64
}

// map of address to Account
type Accounts map[string] Account

func init() {
	rand.Seed(time.Now().UnixNano())
}

// creates a new block on BlockChain
func (bc *BlockChain) NewBlock (nonce uint64, hashPrevBlock string, transactions []Transaction) *Block {

	indexCount := len(bc.chain)
	timestamp := time.Time{}

	block := Block {
		index: indexCount,
		nonce: nonce,
		transactions: transactions,
		timestamp: timestamp,
		hashPrevBlock: hashPrevBlock,
		}

	bc.txQueue = nil
	bc.chain = append(bc.chain, block)
	return &block
}




// proof of work algorithm
func (bc *BlockChain) ProofOfWork (input string) int {

	nonce := 0
	for {
		if bc.ValidProof(input, nonce) {
			return nonce
		}
	nonce += 1
	}
}

// validate proof
func (bc *BlockChain) ValidProof(input string, proof int) bool {

   pfx := []byte(input + string(proof))
   h := sha256.New()
   val := binary.BigEndian.Uint64(h.Sum(pfx))
   return val < rk.DIFFICULTY

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


// Create New Account
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








