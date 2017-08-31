package core

import (
	"fmt"
	"reflect"
	"time"

	"github.com/borisding1994/hathcoin/utils"
)

// TransactionsQueue
type TransactionsQueue chan *Transaction

// BlocksQueue
type BlocksQueue chan Block

// Blockchain
type Blockchain struct {
	CurrentBlock Block
	BlockSlice
	TransactionsQueue
	BlocksQueue
}

func SetupBlockchain() *Blockchain {
	bl := new(Blockchain)
	bl.TransactionsQueue, bl.BlocksQueue = make(TransactionsQueue), make(BlocksQueue)
	//Read blockchain from file and stuff...
	bl.CurrentBlock = bl.CreateNewBlock()
	return bl
}

func (bl *Blockchain) CreateNewBlock() Block {
	prevBlock := bl.BlockSlice.PrevBlock()
	prevBlockHash := []byte{}
	if prevBlock != nil {
		prevBlockHash = prevBlock.CalcHash()
	}
	b := NewBlock(prevBlockHash)
	b.BlockHeader.Origin = currentPeer.Keypair.PublicKey
	return b
}

func (bl *Blockchain) AddBlock(b Block) {
	bl.BlockSlice = append(bl.BlockSlice, b)
}

func DiffTransactionSlices(a, b TransactionSlice) (diff TransactionSlice) {
	//Assumes transaction arrays are sorted (which maybe is too big of an assumption)
	lastJ := 0
	for _, t := range a {
		found := false
		for j := lastJ; j < len(b); j++ {
			if reflect.DeepEqual(b[j].Hash, t.Hash) {
				found = true
				lastJ = j
				break
			}
		}
		if !found {
			diff = append(diff, t)
		}
	}
	return
}

func (bl *Blockchain) GenerateBlocks() chan Block {
	interrupt := make(chan Block)
	go func() {
		block := <-interrupt
	loop:
		utils.Logger.Info("Starting Proof of Work...")
		block.BlockHeader.MerkelRoot = block.GenerateMerkelRoot()
		block.BlockHeader.Nonce = 0
		block.BlockHeader.Timestamp = uint32(time.Now().Unix())
		for true {
			sleepTime := time.Nanosecond
			if block.TransactionSlice.Len() > 0 {
				if CheckProofOfWork(BlockPow, block.CalcHash()) {
					block.Hash = block.Sign(currentPeer.Keypair)
					bl.BlocksQueue <- block
					sleepTime = time.Hour * 24
					fmt.Println("Found Block!")
				} else {
					block.BlockHeader.Nonce += 1
				}

			} else {
				sleepTime = time.Hour * 24
				fmt.Println("No trans sleep")
			}
			select {
			case block = <-interrupt:
				goto loop
			case <-utils.Timeout(sleepTime):
				continue
			}
		}
	}()

	return interrupt
}

func (bl *Blockchain) Run() {

	interruptBlockGen := bl.GenerateBlocks()
	for {
		select {
		case tr := <-bl.TransactionsQueue:

			if bl.CurrentBlock.TransactionSlice.Exists(*tr) {
				continue
			}
			if !tr.VerifyTransaction(TransactionPow) {
				fmt.Println("Recieved non valid transaction", tr)
				continue
			}

			bl.CurrentBlock.AddTransaction(tr)
			interruptBlockGen <- bl.CurrentBlock

			//Broadcast transaction to the network
			mes := NewMessage(MessageSendTransaction)
			mes.Data, _ = tr.MarshalBinary()

			time.Sleep(300 * time.Millisecond)
			currentPeer.Network.BroadcastQueue <- *mes

		case b := <-bl.BlocksQueue:

			if bl.BlockSlice.Exists(b) {
				fmt.Println("block exists")
				continue
			}

			if !b.VerifyBlock(BlockPow) {
				fmt.Println("block verification fails")
				continue
			}

			if reflect.DeepEqual(b.PrevHash, bl.CurrentBlock.CalcHash()) {
				// I'm missing some blocks in the middle. Request'em.
				fmt.Println("Missing blocks in between")
			} else {

				fmt.Println("New block!", b.CalcHash())

				transDiff := TransactionSlice{}

				if !reflect.DeepEqual(b.BlockHeader.MerkelRoot, bl.CurrentBlock.MerkelRoot) {
					// Transactions are different
					fmt.Println("Transactions are different. finding diff")
					transDiff = DiffTransactionSlices(*bl.CurrentBlock.TransactionSlice, *b.TransactionSlice)
				}

				bl.AddBlock(b)

				//Broadcast block and shit
				mes := NewMessage(MessageSendBlock)
				mes.Data, _ = b.MarshalBinary()
				currentPeer.Network.BroadcastQueue <- *mes

				//New Block
				bl.CurrentBlock = bl.CreateNewBlock()
				bl.CurrentBlock.TransactionSlice = &transDiff

				interruptBlockGen <- bl.CurrentBlock
			}
		}
	}
}
