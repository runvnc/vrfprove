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
)

func toVrfPrivKey(phrase string) (crypto.VrfPrivkey) {
  pk, err := mnemonic.ToPrivateKey(phrase)

  if err != nil { panic("invalid vrf mnemonic: %v") }

  var vrfPK crypto.VrfPrivkey
  copy(vrfPK[:], pk[:])
  return vrfPK
}

type Msg []byte

func (m Msg) ToBeHashed() (protocol.HashID, []byte) {
  id, sum := protocol.HashID(""), sha512.Sum512_256(m[:])
  return id, sum[:]
}

func main() {
  sk := toVrfPrivKey(os.Getenv("VRFPRIV"))
  round, err := strconv.Atoi(os.Args[1])
  seedStr := os.Args[2]

  var block_seed [32]byte
  _, err = base32.StdEncoding.Decode(block_seed[:], []byte(seedStr))
  if err != nil { panic("Invalid block seed") }

  var buff = make([]byte, 32+8)
  binary.BigEndian.PutUint64(buff, uint64(round))
  copy(buff[8:], block_seed[:])

  vrfProof, ok := sk.Prove(Msg(buff[:]))
  if !ok { panic("Proof failed.") }

  vrfHash, _ := vrfProof.Hash()
  
  fmt.Print(base32.StdEncoding.EncodeToString(vrfProof[:]))
  fmt.Println(base32.StdEncoding.EncodeToString(vrfHash[:]))
}

