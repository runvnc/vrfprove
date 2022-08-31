package main

import (
  "crypto/sha512"
  "encoding/binary"
  "os"
  "encoding/base32"
  "strconv"
  "fmt"
  "github.com/algorand/go-algorand/protocol"
  "github.com/algorand/go-algorand/crypto"
  "github.com/algorand/go-algorand-sdk/mnemonic"
  "github.com/algorand/go-algorand-sdk/types"
)

func toVrfPrivKey(phrase string) (crypto.VrfPrivkey) {
  pk, err := mnemonic.ToPrivateKey(phrase)
  if err != nil { panic("invalid vrf mnemonic: %v") }

  var vrfPK crypto.VrfPrivkey
  copy(vrfPK[:], pk[:])
  return vrfPK
}

func getPublicKey(s string) (crypto.VrfPubkey) {
  //var key [32]byte
  s = "SJBLLZUWHP6FP27NK47CRZM33ANIDNPUWZIAB3ZGMPD4GEIBHKVVPXMVBQ"
  fmt.Printf("convert public key to byte: [%s] length: %d\n",s, len(s))
  //addr, err := base32.StdEncoding.DecodeString(s)
  key, err := types.DecodeAddress(s)
  
  if err != nil { 
    fmt.Println(err)
    panic("Invalid VRF public key: ")
  }

  var vrfPub crypto.VrfPubkey
  copy(vrfPub[:], key[:])
  return vrfPub
}


type Msg []byte

func (m Msg) ToBeHashed() (protocol.HashID, []byte) {
  id, sum := protocol.HashID(""), sha512.Sum512_256(m[:])
  return id, sum[:]
}

func getRoundSeedHashable(roundStr, seedStr string) (Msg) {
  var block_seed [32]byte
  var buff = make([]byte, 32+8)
 
  round, err := strconv.Atoi(roundStr)
  binary.BigEndian.PutUint64(buff, uint64(round))

  _, err = base32.StdEncoding.Decode(block_seed[:], []byte(seedStr))
  if err != nil { panic("Invalid block seed") }

  copy(buff[8:], block_seed[:])

  return Msg(buff[:])
}

func main() {
  vrfPrivKey := toVrfPrivKey(os.Getenv("VRFPRIV"))

  msg := getRoundSeedHashable(os.Args[2], os.Args[3])

  vrfProof, ok := vrfPrivKey.Prove(msg)

  if !ok { panic("Proof failed.") }

  vrfHash, _ := vrfProof.Hash()
 
	ok1, output := getPublicKey(os.Args[1]).Verify(vrfProof, Msg(msg))
  if !ok1 { panic("Verification failed.") }
 
  fmt.Println(base32.StdEncoding.EncodeToString(vrfProof[:]),
              base32.StdEncoding.EncodeToString(vrfHash[:]),
              base32.StdEncoding.EncodeToString(output[:]))
}

