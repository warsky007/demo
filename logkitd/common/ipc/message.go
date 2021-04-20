package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	MessageHeaderLength  = 8
	LengthFieldOffset    = 0
	LengthFieldLength    = 4
	MessageIdFieldOffset = 4
	MessageIdFieldLength = 4
	MessageContentOffset = 8
)

// | length | message id | content |
// length and message id is 4 bytes
type Message struct {
	MsgId uint32 //
	Data  []byte // message data
	Err   error  // details of any error
}

func BytesToUint32(b []byte) uint32 {
	var value uint32

	binary.Read(bytes.NewReader(b[:4]), binary.BigEndian, &value)

	return value

}

func Encode(message *Message) []byte {
	tmp := new(bytes.Buffer)
	header := make([]byte, MessageHeaderLength)

	contentLength := len(message.Data)
	binary.BigEndian.PutUint32(header, uint32(contentLength+MessageHeaderLength))
	binary.BigEndian.PutUint32(header[MessageIdFieldOffset:], message.MsgId)

	tmp.Write(header)
	tmp.Write(message.Data)
	return tmp.Bytes()
}

func Decode(data []byte) (*Message, error) {
	if len(data) < MessageHeaderLength {
		return nil, fmt.Errorf("data length %d is not enough", len(data))
	}

	length := BytesToUint32(data)
	if len(data) != int(length) {
		return nil, fmt.Errorf("data length %d does not match %d", len(data), length)
	}

	return &Message{
		MsgId: BytesToUint32(data[MessageIdFieldOffset:MessageHeaderLength]),
		Data:  data[MessageContentOffset:],
	}, nil
}
