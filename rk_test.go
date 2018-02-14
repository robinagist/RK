package rk

import "testing"
import (
    "rk/pkg"
    "fmt"
    "time"
)


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
    account := rk.NewAccount()
    if account == nil {
        t.Errorf("no Account created")
    }
}


// create a transaction
func TestCreateTransaction(t *testing.T) {
    sender := rk.NewAccount()
    recipient := rk.NewAccount()
    b := []byte("hello")
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
    bc := new(rk.BlockChain)
    timestamp := time.Now().String()
    blk := bc.NewBlock(0,"", timestamp, nil)
    bc.AddBlock(blk)

    if bc.Size() != 1 {
        t.Errorf("chain did not add Block Zero")
    }

    // find a block
    blk2 := rk.FindBlock(tp, 1)
}
