package core

import (
	"bytes"
	"errors"
	"math"

	"github.com/borisding1994/hathcoin/utils"
)

const (
	MessageTypeSize    = 1
	MessageOptionsSize = 4

	MESSAGE_GET_NODES = iota + 20
	MESSAGE_SEND_NODES

	MessageGetTransaction
	MessageSendTransaction

	MessageGetBlock
	MessageSendBlock
)

type Message struct {
	Identifier byte
	Options    []byte
	Data       []byte

	Reply chan Message
}

func NewMessage(id byte) *Message {

	return &Message{Identifier: id}
}

func (m *Message) MarshalBinary() ([]byte, error) {

	buf := new(bytes.Buffer)

	buf.WriteByte(m.Identifier)
	buf.Write(utils.FitBytesInto(m.Options, MessageOptionsSize))
	buf.Write(m.Data)

	return buf.Bytes(), nil

}

func (m *Message) UnmarshalBinary(d []byte) error {

	buf := bytes.NewBuffer(d)

	if len(d) < MessageOptionsSize+MessageTypeSize {
		return errors.New("Insuficient message size")
	}
	m.Identifier = buf.Next(1)[0]
	m.Options = utils.StripByte(buf.Next(MessageOptionsSize), 0)
	m.Data = buf.Next(math.MaxInt32)

	return nil
}
