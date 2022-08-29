package main

import (
	"crypto/sha512"
	"encoding/binary"
	"os"
	//"strconv"
	//"strings"
	"fmt"

	//"github.com/algorand/go-algorand-sdk/abi"
	//"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	//"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
)

func main() {
	fmt.Println("Started.")
	var pubKeyStr = os.Getenv("VRFPUB")
	var pubKey, _ = types.DecodeAddress(pubKeyStr)
	var vrfMnemonic = os.Getenv("VRFPRIV")

	var vrfPrivateKey, err = mnemonic.ToPrivateKey(vrfMnemonic)

	if err != nil {
		fmt.Printf("invalid vrf mnemonic: %v", err)
		return
	}

	var vrfp = NewVRF(pubKey[:], vrfPrivateKey)
  var block_seed = 23452345234533453
  var round = 83483452

	var buff = make([]byte, 32+8)
	binary.BigEndian.PutUint64(buff, uint64(round))
	copy(buff[8:], block_seed[:])
	var vrfInput = sha512.Sum512_256(buff[:])

	var hash, proof = vrfp.Prove(vrfInput[:])
     
	fmt.Print(hash)
	fmt.Print(proof)
}
/*

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
*/

