package rk

import "testing"
import (
    "rk/pkg"
    "fmt"
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

    fmt.Println("sender tx", sender)
    fmt.Println("transaction", nta)

    nta, err = rk.NewTransaction(sender, recipient, 0, 0, b)

    if err != nil {
        fmt.Println(err)
        t.Errorf("unable to create new transaction")
    }
    fmt.Println("sender tx", sender)

    // create a block


}
