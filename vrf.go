package main

import (
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/protocol"
)

type AlgorandVRF struct {
	PublicKey  crypto.VrfPubkey
	PrivateKey crypto.VrfPrivkey
}

func NewVRF(publicKey, privateKey []byte) *AlgorandVRF {
	privKey := [64]byte{}
	copy(privKey[:], privateKey[:])

	pubKey := [32]byte{}
	copy(pubKey[:], publicKey[:])

	return &AlgorandVRF {
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}
}

type Msg []byte

func (m Msg) ToBeHashed() (protocol.HashID, []byte) {
	return protocol.HashID(""), m[:]
}

func (avp *AlgorandVRF) Prove(msg []byte) (vrf []byte, proof []byte) {
	vrfProof, _ := avp.PrivateKey.Prove(Msg(msg))
	vrfHash, _ := vrfProof.Hash()
	return vrfHash[:], vrfProof[:]
}

func (avp *AlgorandVRF) Verify(msg, proof []byte) bool {
	var vrfProof crypto.VrfProof
	copy(vrfProof[:], proof[:])
	ok, _ := avp.PublicKey.Verify(vrfProof, Msg(msg))
	return ok
}
