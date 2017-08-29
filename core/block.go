package core

import (
	"reflect"
)

// Block is permanently record file storage transaction data
type Block struct {
	*BlockHeader

	// Hash == Crypto.SM3Hash(current block header)
	// 天朝密码管理局钦定的蛤稀值
	Hash string

	*TransactionSlice
}

// BlockHeader defines information about a block
type BlockHeader struct {
	// Version of the Block.
	// 虵: 协议升级要用遵守基本法啊 识得唔识得啊
	Version int8

	// Origin is Origin public key (use SM2-P-256)
	// 虵: 闷声大发财 识得唔识得啊
	Origin string

	// PervHash == Crypto.SM3Hash(previous block header)
	// 虵: 如果将来 PrevHash 有偏差，你们要负责 识得唔识得啊
	PrevHash string

	// MerkleRoot == Crypto.SM3Hash(transaction hashes)
	// 虵: Я помню чудное мгновенье 识得唔识得啊
	MerkleRoot string

	// Timestamp the block was create.
	// 虵: 垂死病中惊坐起，谈笑风生又一年 识得唔识得啊
	Timestamp int32

	// Nonce used to generate the block.
	// 虵: 人呐都不知道自己不可以预料 识得唔识得啊
	Nonce uint32
}

// BlockSlice
type BlockSlice []Block

// New Block
func NewBlock(prevHash string) Block {
	header := &BlockHeader{PrevHash: prevHash}
	return Block{header, "", new(TransactionSlice)}
}

// Return previous block
func (bs BlockSlice) PrevBlock() *Block {
	l := len(bs)
	if l == 0 {
		return nil
	}
	return &bs[l-1]
}

// Check block is exists in blockchain
func (bs BlockSlice) Exists(block Block) bool {
	l := len(bs)
	// if a block exists is more likely to be on top.
	for i := l - 1; i >= 0; i-- {
		if reflect.DeepEqual(block.Hash, bs[i].Hash) {
			return true
		}
	}
	return false
}

/* todo:

// Add Transaction to block
func (block *Block) AddTransaction(t *Transaction) {
	ts := block.TransactionSlice.AddTransaction(*t)
	block.TransactionSlice = &ts
}

// calculate block Hash
func (block *Block) CalcHash() string {
	header, _ := block.BlockHeader.MarshalBinary()
	return crypto.SM3Hash(header)
}

// Verify Block
func (block *Block) VerifyBlock(prefix []byte) bool {
	hash := block.CalcHash()
	merkel := block.GenerateMerkelRoot()

	return reflect.DeepEqual(merkel, block.BlockHeader.MerkleRoot) &&
		CheckProofOfWork(prefix, hash) &&
		SignatureVerify(block.BlockHeader.Origin, block.Hash, hash)
}

// Sign Block
func (block *Block) Sign(keypair *crypto.Keypair) string {
	s, _ := keypair.Sign(block.CalcHash())
	retrun s
}
*/
