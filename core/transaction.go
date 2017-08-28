package core

// Transaction is a transfer of HathCoin value that is broadcast to the network and collected into blocks.
type Transaction struct {
	Header TransactionHeader

	// Hash == Crypto.SM3Hash(current transaction header)
	Hash string

	// Payload is raw transaction data.
	Payload []byte
}

// TransactionHeader defines information about a transaction
type TransactionHeader struct {
	// From is origin public key
	From string

	// To is destination public key
	To string

	// Timestamp the transaction was create.
	Timestamp int32

	// PayloadHash == Crypto.SM3Hash(current transaction payload)
	PayloadHash string

	// PayloadLength == len(current transaction payload)
	PayloadLength uint32

	// Nonce Proof of work.
	Nonce uint32
}

// TransactionSlice
type TransactionSlice []Transaction
