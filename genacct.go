package main

import (
  "fmt"
  "github.com/algorand/go-algorand-sdk/mnemonic"
  "github.com/algorand/go-algorand-sdk/crypto"
)

func main() {
  var acct = crypto.GenerateAccount()
  fmt.Println(acct.Address)
  var mnemonic, _ = mnemonic.FromPrivateKey(acct.PrivateKey) 
  fmt.Println(mnemonic)
}

