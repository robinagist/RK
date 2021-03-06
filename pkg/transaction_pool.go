package rk

import (
	"github.com/cbergoon/merkletree"
	"encoding/json"
	"crypto/sha256"
	"encoding/hex"
)

// Transaction Pool
type TransactionPool struct {
	tPool []Transaction
}

func (tp *TransactionPool) Filter(criteria string) []Transaction {
	// TODO make this actually filter on criteria
	return tp.tPool
}

func (tp *TransactionPool) Add(tx *Transaction) bool {
	// make sure the same transaction is not being added twice -- if so, reject
	tp.tPool = append(tp.tPool, *tx)
	return true
}

func (tp *TransactionPool) Size() int {
	return len(tp.tPool)
}

type Content struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t Content) CalculateHash() []byte {
	h := sha256.New()
	h.Write([]byte(t.x))
	return h.Sum(nil)
}

//Equals tests for equality of two Contents
func (t Content) Equals(other merkletree.Content) bool {
	return t.x == other.(Content).x
}

// generates the merkle root
func MerkleRoot(txs []Transaction) string {
	var list []merkletree.Content
	for _, tx := range txs {
		mtx, err := json.Marshal(tx)
		if err != nil {
            panic(err)
		}
		mtc := Content{x: string(mtx)}
		list = append(list, mtc)
	}

	tree, _ := merkletree.NewTree(list)
	root := tree.MerkleRoot()
	return hex.EncodeToString(root)
}