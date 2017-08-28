package core

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
