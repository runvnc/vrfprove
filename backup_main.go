package main

import (
	"crypto/sha512"
	"encoding/binary"
	"os"
  "math/rand"
  "bufio"
	//"strconv"
	//"strings"
	"fmt"

	"github.com/algorand/go-algorand/protocol"
	"github.com/algorand/go-algorand/crypto"
	//"github.com/algorand/go-algorand-sdk/abi"
	//"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/mnemonic"
	//"github.com/algorand/go-algorand-sdk/crypto"
//	"github.com/algorand/go-algorand-sdk/types"
)

func toVrfPrivKey(phrase string) (crypto.VrfPrivkey) {
	pk, err := mnemonic.ToPrivateKey(phrase)

  if err != nil {
		panic("invalid vrf mnemonic: %v")
	}

  var vrfPK crypto.VrfPrivkey
  copy(vrfPK[:], pk[:])
  return vrfPK
}

type Msg []byte

//return protocol.HashID(""), m[:]

func (m Msg) ToBeHashed() (protocol.HashID, []byte) {
  var vrfInput = sha512.Sum512_256(m)


	return protocol.HashID(""), m[:]
}

func main() {
	sk := toVrfPrivKey(os.Getenv("VRFPRIV"))
	var stdin = bufio.NewReader(os.Stdin)


	var block_seed [32]byte
  for i := 0; i < 32; i++ {
    block_seed[i] = byte(rand.Intn(255))
  }

  var round = 834834520

	var buff = make([]byte, 32+8)
	binary.BigEndian.PutUint64(buff, uint64(round))
	copy(buff[8:], block_seed[:])
	var vrfInput = sha512.Sum512_256(buff[:])

	vrfProof, _ := sk.Prove(Msg(vrfInput[:]))
	vrfHash, _ := vrfProof.Hash()
     
	fmt.Println(vrfProof)
  fmt.Println("--------------------")
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

