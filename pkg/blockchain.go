package rk

import (
	"time"
	"encoding/json"
	"crypto/sha256"
	"math/rand"
	"encoding/binary"
	"rk/internal"
	"errors"
	"strconv"
	"fmt"
	"encoding/hex"
)


type Block struct {
	Index         int        `json:"index"`
	Nonce         uint64     `json:"nonce"`
	Timestamp     string     `json:"timestamp"`
	Transactions  []Transaction    `json:"transactions"`
	HashPrevBlock string     `json:"hashPrevBlock"`
	MerkleRootHash string    `json:"merkleRootHash"`
}


type BlockChain struct {
    chain []Block
    target uint64
}


func (bc *BlockChain) init() {
	rand.Seed(time.Now().UnixNano())
}

// creates the first block on BlockChain
func (bc *BlockChain) GenerateZeroBlock () *Block {

	indexCount := len(bc.chain)
	if indexCount != 0 {
		errors.New("chain has already been initialized")
	}
    ts := time.Now().String()

	block := Block {
		Index: 0,
		Nonce: 0,
		Timestamp: ts,
		Transactions: nil,
		HashPrevBlock: "1rk.block.zero",
		MerkleRootHash: "noor",
		}
	return &block
}

func (bc *BlockChain) AddBlock (block *Block) {
	bc.chain = append(bc.chain, *block)
}

// returns the hash of the previous block
func (bc *BlockChain) hashPrevBlock() string {
	block := bc.chain[len(bc.chain)-1]
	ks := string(Hash(&block))
	fmt.Println(ks)
	return ks
}


// size of the chain
func (bc *BlockChain) Size() int {
	return len(bc.chain)
}

func (bc *BlockChain) FindBlock(tp *TransactionPool) *Block {

	bc.init()
	block := new(Block)
	block.Index = bc.Size()
	block.Timestamp = time.Now().String()
	block.HashPrevBlock = bc.hashPrevBlock()

	// get the desired transactions from the transaction pool
	block.Transactions = tp.Filter("")
	block.MerkleRootHash = MerkleRoot(block.Transactions)

	// proof of work
	sequence := block.HashPrevBlock + block.MerkleRootHash
	block.Nonce = ProofOfWork(sequence)
	return block

}


// hash block
func Hash(block *Block) string {

	blockJson, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}
	// get a hex digest
	h := sha256.New()
	h.Write(blockJson)
	return hex.EncodeToString(h.Sum(nil))
}


// proof of work
func ProofOfWork (input string) uint64 {
	var nonce uint64
	nonce = 0
	for {
		if ValidProof(input, nonce) {
			return nonce
		}
		nonce += 1
	}
}

// validate proof
func ValidProof(inp string, proof uint64) bool {

	pfx := inp + strconv.Itoa(int(proof))
	pfxx := []byte(pfx)
	h := sha256.New()
	h.Write(pfxx)
	val := binary.BigEndian.Uint64(h.Sum(nil))

	return val < rk.TARGET

}






