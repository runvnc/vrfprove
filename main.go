package main

import (
	"context"
	"crypto/sha512"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"fmt"

	"github.com/barnjamin/vrf-oracle/sandbox"
	"github.com/barnjamin/vrf-oracle/vrfproducers/algorand"

	"github.com/algorand/go-algorand-sdk/abi"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/algorand/go-algorand/rpcs"
	"github.com/algorand/indexer/fetcher"
	"github.com/sirupsen/logrus"
)

var (
	vrfp        *AlgorandVRF
	accts    []crypto.Account
)

func main() {
	fmt.Println("Started.")

	vrfPrivateKey, err := mnemonic.ToPrivateKey(vrfMnemonicString)
	if err != nil {
		fmt.Printf("invalid vrf mnemonic: %v", err)
		return
	}

	vrfp = NewVRF(accts[0].PublicKey[:], accts[0].PrivateKey[:])

	var block_seed = cert.Block.Seed()
	var round = cert.Block.Round()

	buff := make([]byte, 32+8)
	binary.BigEndian.PutUint64(buff, uint64(round))
	copy(buff[8:], block_seed[:])
	vrfInput := sha512.Sum512_256(buff[:])

	hash, proof := vrfp.Prove(vrfInput[:])
     
	fmt.Print(hash)
	fmt.Print(proof)
}


