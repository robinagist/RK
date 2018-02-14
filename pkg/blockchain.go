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
	timestamp     string
	hashPrevBlock string
	merkleRootHash string
}


type BlockChain struct {
    chain []Block
    target uint64
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

// creates a new block on BlockChain
func (bc *BlockChain) NewBlock (nonce uint64, hashPrevBlock string, ts string, transactions []Transaction) *Block {

	indexCount := len(bc.chain)

	block := Block {
		index: indexCount,
		nonce: nonce,
		timestamp: ts,
		hashPrevBlock: hashPrevBlock,
		merkleRootHash: "",
		}

	return &block
}

func (bc *BlockChain) AddBlock (block *Block) {
	bc.chain = append(bc.chain, *block)
}

// returns the hash of the previous block
func (bc *BlockChain) hashPrevBlock() string {
	block := bc.chain[len(bc.chain)-1]
	return string(Hash(&block))
}


// size of the chain
func (bc *BlockChain) Size() int {
	return len(bc.chain)
}

func (bc *BlockChain) FindBlock(tp *TransactionPool) *Block {

	block := new(Block)
	block.index = bc.Size()
	block.timestamp = time.Now().String()
	block.hashPrevBlock = bc.hashPrevBlock()

	// get the desired transactions from the transaction pool
	transactions := tp.Filter("")
	block.merkleRootHash = MerkelRoot(transactions)

	// proof of work
	sequence := block.timestamp + block.hashPrevBlock + block.merkleRootHash
	block.nonce = ProofOfWork(sequence)
	return block

}


func GenerateAddressString() string {
	t := string(time.Now().Nanosecond())
	n := sha3.New256()
	n.Write([]byte(t + string(rand.Intn(1000000))))
	return "099" + hex.EncodeToString(n.Sum(nil))
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
func ValidProof(input string, proof uint64) bool {

	pfx := []byte(input + string(proof))
	h := sha256.New()
	val := binary.BigEndian.Uint64(h.Sum(pfx))
	return val < rk.DIFFICULTY

}






