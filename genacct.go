package main

import (
	//"crypto/sha512"
	//"encoding/binary"
	//"os"
	//"strconv"
	//"strings"
	"fmt"

	//"github.com/algorand/go-algorand-sdk/abi"
	//"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/crypto"
	//"github.com/algorand/go-algorand-sdk/types"
)

func main() {
	var acct = crypto.GenerateAccount()
  fmt.Println(acct.Address)
  var mnemonic, _ = mnemonic.FromPrivateKey(acct.PrivateKey) 
  fmt.Println(mnemonic)
}

