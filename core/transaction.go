package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"reflect"
	"time"

	"github.com/borisding1994/hathcoin/utils"
	"github.com/borisding1994/hathcoin/utils/crypto"
)

// Transaction is a transfer of HathCoin value that is broadcast to the network and collected into blocks.
type Transaction struct {
	Header TransactionHeader

	// Hash == Crypto.SM3Hash(current transaction header)
	Hash []byte

	// Payload is raw transaction data.
	Payload []byte
}

// TransactionHeader defines information about a transaction
type TransactionHeader struct {
	// From is origin public key
	From []byte

	// To is destination public key
	To []byte

	// Timestamp the transaction was create.
	Timestamp uint32

	// PayloadHash == Crypto.SM3Hash(current transaction payload)
	PayloadHash []byte

	// PayloadLength == len(current transaction payload)
	PayloadLength uint32

	// Nonce Proof of work.
	Nonce uint32
}

// TransactionSlice
type TransactionSlice []Transaction

// Returns bytes to be sent to the network
func NewTransaction(from, to, payload []byte) *Transaction {

	t := Transaction{Header: TransactionHeader{From: from, To: to}, Payload: payload}

	t.Header.Timestamp = uint32(time.Now().Unix())
	t.Header.PayloadHash = []byte(crypto.SM3HashByte(t.Payload))
	t.Header.PayloadLength = uint32(len(t.Payload))

	return &t
}

// CalcHash can calculate Transaction Hash
func (t *Transaction) CalcHash() []byte {
	h, _ := t.Header.MarshalBinary()
	return crypto.SM3HashByte(h)
}

// Sign Transaction
func (t *Transaction) Sign(keypair *crypto.Keypair) []byte {
	s, err := keypair.Sign(t.CalcHash())
	if err != nil {
		utils.Logger.Error("Sign Block Error", err)
	}
	return s
}

// VerifyTransaction by Hash
func (t *Transaction) VerifyTransaction(pow []byte) bool {
	headerHash := t.CalcHash()
	payloadHash := crypto.SM3HashByte(t.Payload)

	return reflect.DeepEqual(payloadHash, t.Header.PayloadHash) &&
		CheckProofOfWork(pow, headerHash) &&
		crypto.VerifySign(t.Header.From, t.Hash, headerHash)
}

// GenerateNonce to TransactionHeader
func (t *Transaction) GenerateNonce(prefix []byte) uint32 {
	newT := t
	for {
		if CheckProofOfWork(prefix, newT.CalcHash()) {
			break
		}
		newT.Header.Nonce++
	}
	return newT.Header.Nonce
}

func (t *Transaction) MarshalBinary() ([]byte, error) {
	headerBytes, _ := t.Header.MarshalBinary()
	if len(headerBytes) != TransactionHeaderSize {
		return nil, errors.New("Header marshalling error")
	}
	return append(append(headerBytes, utils.FitBytesInto(t.Hash, Sm2PublicKeySize)...), t.Payload...), nil
}

func (t *Transaction) UnmarshalBinary(d []byte) ([]byte, error) {
	buf := bytes.NewBuffer(d)
	if len(d) < TransactionHeaderSize+Sm2PublicKeySize {
		return nil, errors.New("Insuficient bytes for unmarshalling transaction")
	}
	h := &TransactionHeader{}
	if err := h.UnmarshalBinary(buf.Next(TransactionHeaderSize)); err != nil {
		return nil, err
	}
	t.Header = *h
	t.Hash = utils.StripByte(buf.Next(Sm2PublicKeySize), 0)
	t.Payload = buf.Next(int(t.Header.PayloadLength))
	return buf.Next(math.MaxInt32), nil
}

func (th *TransactionHeader) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(utils.FitBytesInto(th.From, Sm2PublicKeySize))
	buf.Write(utils.FitBytesInto(th.To, Sm2PublicKeySize))
	binary.Write(buf, binary.LittleEndian, th.Timestamp)
	buf.Write(utils.FitBytesInto(th.PayloadHash, 32))
	binary.Write(buf, binary.LittleEndian, th.PayloadLength)
	binary.Write(buf, binary.LittleEndian, th.Nonce)
	return buf.Bytes(), nil
}

func (th *TransactionHeader) UnmarshalBinary(d []byte) error {
	buf := bytes.NewBuffer(d)
	th.From = utils.StripByte(buf.Next(Sm2PublicKeySize), 0)
	th.To = utils.StripByte(buf.Next(Sm2PublicKeySize), 0)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Timestamp)
	th.PayloadHash = buf.Next(32)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.PayloadLength)
	binary.Read(bytes.NewBuffer(buf.Next(4)), binary.LittleEndian, &th.Nonce)
	return nil
}

func (slice TransactionSlice) Len() int {
	return len(slice)
}

func (slice TransactionSlice) Exists(tr Transaction) bool {
	for _, t := range slice {
		if reflect.DeepEqual(t.Hash, tr.Hash) {
			return true
		}
	}
	return false
}

func (slice TransactionSlice) AddTransaction(t Transaction) TransactionSlice {
	// Inserted sorted by timestamp
	for i, tr := range slice {
		if tr.Header.Timestamp >= t.Header.Timestamp {
			return append(append(slice[:i], t), slice[i:]...)
		}
	}
	return append(slice, t)
}

func (slice *TransactionSlice) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, t := range *slice {
		bs, err := t.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(bs)
	}
	return buf.Bytes(), nil
}

func (slice *TransactionSlice) UnmarshalBinary(d []byte) error {
	remaining := d
	for len(remaining) > TransactionHeaderSize+Sm2PublicKeySize {
		t := new(Transaction)
		rem, err := t.UnmarshalBinary(remaining)
		if err != nil {
			return err
		}
		*slice = append(*slice, *t)
		remaining = rem
	}
	return nil
}
