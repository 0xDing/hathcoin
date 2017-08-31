package core

import (
	"bufio"
	"net"
	"os"
	"strings"

	"github.com/borisding1994/hathcoin/config"
	"github.com/borisding1994/hathcoin/utils"
	"github.com/borisding1994/hathcoin/utils/crypto"
)

const (
	Sm2PublicKeySize = 88

	TransactionHeaderSize = Sm2PublicKeySize /* from key */ +
		Sm2PublicKeySize /* to key */ +
		4 /* int32 timestamp */ +
		32 /* SM3 payload hash */ +
		4 /* int32 payload length */ +
		4 /* int32 nonce */

	BlockHeaderSize = Sm2PublicKeySize /* origin key */ +
		4 /* int32 timestamp */ +
		32 /* prev block hash */ +
		32 /* merkel tree hash */ +
		4 /* int32 nonce */

	TransactionPowComplexity     = 2
	TestTransactionPowComplexity = 1

	BlockPowComplexity     = 4
	TestBlockPowComplexity = 2

	PowPrefix = 0

	// Identifies Hathcoin protocol version
	HTCProtoVersion = 1
)

var currentPeer = struct {
	Keypair *crypto.Keypair
	*Blockchain
	*Network
}{}

func Run() {
	LoadKeypair()
	addr := strings.Split(config.GetString("addr"), ":")
	//Setup Network
	currentPeer.Network = SetupNetwork(addr[0], addr[1])
	go currentPeer.Network.Run()
	// find initial peer by DNS SRV
	_, addrs, err := net.LookupSRV("hathcoin-discovery", "tcp", config.GetString("initial_peers"))
	if err != nil {
		utils.Logger.Error(err)
	}
	for _, peerAddr := range addrs {
		currentPeer.Network.ConnectionsQueue <- peerAddr.Target
	}
	currentPeer.Blockchain = SetupBlockchain()
	go currentPeer.Blockchain.Run()

	//Read Stdin to create transations
	stdin := ReadStdin()
	for {
		select {
		case str := <-stdin:
			currentPeer.Blockchain.TransactionsQueue <- CreateTransaction(str)
		case msg := <-currentPeer.Network.IncomingMessages:
			HandleIncomingMessage(msg)
		}
	}
}

func ReadStdin() chan string {
	cb := make(chan string)
	sc := bufio.NewScanner(os.Stdin)
	go func() {
		for sc.Scan() {
			cb <- sc.Text()
		}
	}()
	return cb
}

func HandleIncomingMessage(msg Message) {
	switch msg.Identifier {
	case MessageSendTransaction:
		t := new(Transaction)
		_, err := t.UnmarshalBinary(msg.Data)
		if err != nil {
			utils.Logger.Error(err)
			break
		}
		currentPeer.Blockchain.TransactionsQueue <- t
	case MessageSendBlock:
		b := new(Block)
		err := b.UnmarshalBinary(msg.Data)
		if err != nil {
			utils.Logger.Error(err)
			break
		}
		currentPeer.Blockchain.BlocksQueue <- *b
	}
}

func CreateTransaction(txt string) *Transaction {
	t := NewTransaction(currentPeer.Keypair.PublicKey, nil, []byte(txt))
	t.Header.Nonce = t.GenerateNonce(TransactionPow)
	t.Hash = t.Sign(currentPeer.Keypair)
	return t
}
