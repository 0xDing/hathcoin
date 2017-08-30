package core

import (
	"reflect"

	"github.com/borisding1994/hathcoin/utils"
)

var (
	TransactionPow = utils.ArrayOfBytes(TransactionPowComplexity, PowPrefix)
	BlockPow       = utils.ArrayOfBytes(BlockPowComplexity, PowPrefix)

	TestTransactionPow = utils.ArrayOfBytes(TestTransactionPowComplexity, PowPrefix)
	TestBlockPow       = utils.ArrayOfBytes(TestBlockPowComplexity, PowPrefix)
)

func CheckProofOfWork(prefix []byte, hash []byte) bool {
	if len(prefix) > 0 {
		return reflect.DeepEqual(prefix, hash[:len(prefix)])
	}
	return true
}
