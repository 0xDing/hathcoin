package core

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

func Run() {
	LoadKeypair()
}
