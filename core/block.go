package core

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"

	"github.com/borisding1994/hathcoin/utils"
	"github.com/borisding1994/hathcoin/utils/crypto"
)

// Block is permanently record file storage transaction data
type Block struct {
	*BlockHeader

	// Hash == Crypto.SM3Hash(current block header)
	// 天朝密码管理局钦定的蛤稀值
	Hash []byte

	*TransactionSlice
}

// BlockHeader defines information about a block
type BlockHeader struct {
	// Version of the Block.
	// 虵: 协议升级要用遵守基本法啊 识得唔识得啊
	Version int8

	// Origin is Origin public key (use SM2-P-256)
	// 虵: 闷声大发财 识得唔识得啊
	Origin []byte

	// PervHash == Crypto.SM3Hash(previous block header)
	// 虵: 如果将来 PrevHash 有偏差，你们要负责 识得唔识得啊
	PrevHash []byte

	// MerkelRoot == Crypto.SM3Hash(transaction hashes)
	// 虵: Я помню чудное мгновенье 识得唔识得啊
	MerkelRoot []byte

	// Timestamp the block was create.
	// 虵: 垂死病中惊坐起，谈笑风生又一年 识得唔识得啊
	Timestamp uint32

	// Nonce used to generate the block.
	// 虵: 人呐都不知道自己不可以预料 识得唔识得啊
	Nonce uint32
}

// BlockSlice
type BlockSlice []Block

// New Block
func NewBlock(prevHash []byte) Block {
	header := &BlockHeader{PrevHash: prevHash}
	return Block{header, nil, new(TransactionSlice)}
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

// AddTransaction to block
func (block *Block) AddTransaction(t *Transaction) {
	ts := block.TransactionSlice.AddTransaction(*t)
	block.TransactionSlice = &ts
}

// CalcHash can calculate block Hash
func (block *Block) CalcHash() []byte {
	header, _ := block.BlockHeader.MarshalBinary()
	return crypto.SM3HashByte(header)
}

// Verify Block
func (block *Block) VerifyBlock(prefix []byte) bool {
	hash := block.CalcHash()
	merkel := block.GenerateMerkelRoot()

	return reflect.DeepEqual(merkel, block.BlockHeader.MerkelRoot) &&
		CheckProofOfWork(prefix, hash) &&
		crypto.VerifySign(block.BlockHeader.Origin, block.Hash, hash)
}

// Sign Block
func (block *Block) Sign(keypair *crypto.Keypair) []byte {
	s, err := keypair.Sign(block.CalcHash())
	if err != nil {
		utils.Logger.Error("Sign Block Error", err)
	}
	return s
}

// GenerateNonce to BlockHeader
func (block *Block) GenerateNonce(prefix []byte) uint32 {
	b := block
	for {
		if CheckProofOfWork(prefix, b.CalcHash()) {
			break
		}
		b.BlockHeader.Nonce++
	}
	return b.BlockHeader.Nonce
}

func (block *Block) GenerateMerkelRoot() []byte {
	var merkell func(hashes [][]byte) []byte
	merkell = func(hashes [][]byte) []byte {
		l := len(hashes)
		if l == 0 {
			return nil
		}
		if l == 1 {
			return hashes[0]
		} else {
			if l%2 == 1 {
				return merkell([][]byte{merkell(hashes[:l-1]), hashes[l-1]})
			}
			bs := make([][]byte, l/2)
			for i := range bs {
				j, k := i*2, (i*2)+1
				bs[i] = crypto.SM3HashByte(append(hashes[j], hashes[k]...))
			}
			return merkell(bs)
		}
	}
	ts := utils.Map(func(t Transaction) []byte { return t.CalcHash() }, []Transaction(*block.TransactionSlice)).([][]byte)
	return merkell(ts)
}

func (block *Block) MarshalBinary() ([]byte, error) {
	bhb, err := block.BlockHeader.MarshalBinary()
	if err != nil {
		return nil, err
	}
	sig := utils.FitBytesInto(block.Hash, Sm2PublicKeySize)
	tsb, err := block.TransactionSlice.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return append(append(bhb, sig...), tsb...), nil
}

func (block *Block) UnmarshalBinary(d []byte) error {
	buf := bytes.NewBuffer(d)
	header := new(BlockHeader)
	err := header.UnmarshalBinary(buf.Next(BlockHeaderSize))
	if err != nil {
		return err
	}
	block.BlockHeader = header
	block.Hash = utils.StripByte(buf.Next(Sm2PublicKeySize), 0)
	ts := new(TransactionSlice)
	err = ts.UnmarshalBinary(buf.Next(math.MaxInt32))
	if err != nil {
		return err
	}
	block.TransactionSlice = ts
	return nil
}

func (h *BlockHeader) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(utils.FitBytesInto(h.Origin, Sm2PublicKeySize))
	binary.Write(buf, binary.LittleEndian, h.Timestamp)
	buf.Write(utils.FitBytesInto(h.PrevHash, 32))
	buf.Write(utils.FitBytesInto(h.MerkelRoot, 32))
	binary.Write(buf, binary.LittleEndian, h.Nonce)
	return buf.Bytes(), nil
}

func (h *BlockHeader) UnmarshalBinary(d []byte) error {
	buf := bytes.NewBuffer(d)
	h.Origin = utils.StripByte(buf.Next(Sm2PublicKeySize), 0)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &h.Timestamp)
	h.PrevHash = buf.Next(32)
	h.MerkelRoot = buf.Next(32)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &h.Nonce)
	return nil
}
