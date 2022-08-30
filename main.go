package main

import (
//	"crypto/sha512"
	"encoding/binary"
	"os"
  "math/rand"
	//"strconv"
	//"strings"
	"fmt"
  "reflect"

	"github.com/algorand/go-algorand/crypto"
	//"github.com/algorand/go-algorand-sdk/abi"
	//"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	//"github.com/algorand/go-algorand-sdk/crypto"
//	"github.com/algorand/go-algorand-sdk/types"
)

func main() {
//	var pubKeyStr = os.Getenv("VRFPUB")
//	var pubKey, _ = types.DecodeAddress(pubKeyStr)
	var vrfMnemonic = os.Getenv("VRFPRIV")

	privKey, err := mnemonic.ToPrivateKey(vrfMnemonic)

	privKey_ := [64]byte{}
	copy(privKey_[:], privKey[:])

	//alpha = make([]byte, hex.DecodedLen(len(alphaHex)))
	//mustDecode(t, alpha, alphaHex)

	//pkTest, sk := VrfKeygenFromSeed(seed)

	//piTest, ok := sk.proveBytes(alpha)
	
  var vrfPrivateKeyx crypto.VrfPrivkey
  copy(vrfPrivateKeyx[:], privKey_[:])

	if err != nil {
		fmt.Printf("invalid vrf mnemonic: %v", err)
		return
	}

	var block_seed [32]byte
  for i := 0; i < 32; i++ {
    block_seed[i] = byte(rand.Intn(255))
  }

  var round = 834834520

	var buff = make([]byte, 32+8)
	binary.BigEndian.PutUint64(buff, uint64(round))
	copy(buff[8:], block_seed[:])
	//var vrfInput = sha512.Sum512_256(buff[:])

  fooType := reflect.TypeOf(vrfPrivateKeyx)
  for i := 0; i < fooType.NumMethod(); i++ {
    method := fooType.Method(i)
    fmt.Println(method.Name)
  }

	vrfProof, _ := vrfPrivateKeyx.Prove(vrfInput[:])
	vrfHash, _ := vrfProof.Hash()
     
	fmt.Print(vrfProof)
	fmt.Print(vrfHash)
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

