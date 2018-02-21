package rk

import "testing"
import (
    "rk/pkg"
    "fmt"
)


var BC *rk.BlockChain

func TestTrue(t *testing.T) {

}

// accounts

func TestGenerateAddress(t *testing.T) {
    addr := rk.GenerateAddressString()
    if addr == "" {
        t.Errorf("address not generated")
    }
    fmt.Println(addr)
}

func TestCreateAccount(t *testing.T) {
    a := new(rk.Accounts)
    account := a.NewAccount("")
    if account == nil {
        t.Errorf("no Account created")
    }
}


// create a transaction
func TestCreateTransaction(t *testing.T) {

    accounts := new(rk.Accounts)

    sender := accounts.NewAccount("")
    recipient := accounts.NewAccount("")
    b := "hello world"
    nta, err := rk.NewTransaction(sender, recipient, 0, 0, b)

    if err != nil {
        fmt.Println(err)
        t.Errorf("error")
    }

    // create a transaction pool
    tp := new(rk.TransactionPool)
    tp.Add(nta)
    nta, err = rk.NewTransaction(sender, recipient, 0, 0, b)

    if err != nil {
        fmt.Println(err)
        t.Errorf("unable to create new transaction")
    }

    tp.Add(nta)

    if tp.Size() != 2 {
        t.Errorf("size mismatch:  should be 2 - got ", tp.Size())
    }

    // create a block
    // start with block zero
    BC := new(rk.BlockChain)
 //   timestamp := time.Now().String()
    blk := BC.GenerateZeroBlock()
    BC.AddBlock(blk)

    if BC.Size() != 1 {
        t.Errorf("chain did not add zero block")
    } else {fmt.Println("chain added zero block")}

    // find a block
    blk2 := BC.FindBlock(tp)
    BC.AddBlock(blk2)
    fmt.Println(blk2.Nonce)

    sequence := blk2.HashPrevBlock + blk2.MerkleRootHash
    if !rk.ValidProof(sequence, blk2.Nonce) {
        t.Errorf("Nonce does not validate")
    } else {fmt.Println("nonce validates")}

    // find another block without dumping transaction pool
    blk3 := BC.FindBlock(tp)
    BC.AddBlock(blk3)
    fmt.Println(blk3.Nonce)

    if BC.Size() != 3 {
        t.Errorf("chain should have 3 blocks - only has ", BC.Size())
    }
    sequence = blk3.HashPrevBlock + blk3.MerkleRootHash
    if !rk.ValidProof(sequence, blk3.Nonce) {
        t.Errorf("Nonce does not validate")
    } else {fmt.Println("nonce validates")}


}
