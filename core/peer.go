package core

import (
	"github.com/borisding1994/hathcoin/utils/crypto"
)

var Peer = struct {
	Keypair *crypto.Keypair
	*Blockchain
	//*Network
}{}

func CreateTransaction(txt string) *Transaction {
	t := NewTransaction(Peer.Keypair.PublicKey, nil, []byte(txt))
	t.Header.Nonce = t.GenerateNonce(TransactionPow)
	t.Hash = t.Sign(Peer.Keypair)
	return t
}
