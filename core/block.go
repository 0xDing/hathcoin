package core

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
	// 虵: 闷声大发财
	Origin string

	// PervHash == Crypto.SM3Hash(previous block header)
	// 虵: 如果将来 PrevHash 有偏差，你们要负责
	PrevHash string

	// MerkleRoot == Crypto.SM3Hash(transaction hashes)
	// 虵: Я помню чудное мгновенье
	MerkleRoot string

	// Timestamp the block was create.
	// 虵: 垂死病中惊坐起，谈笑风生又一年
	Timestamp int32

	// Nonce used to generate the block.
	// 虵: 人呐都不知道自己不可以预料
	Nonce uint32
}

// BlockSlice
type BlockSlice []Block
