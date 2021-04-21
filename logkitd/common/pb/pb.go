package pb

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Pb struct{}

func NewPb() *Pb {
	return &Pb{}
}

func (pb *Pb) Marshal(v interface{}) ([]byte, error) {
	message, ok := v.(proto.Message)
	if ok {
		return proto.Marshal(message)
	} else {
		return nil, fmt.Errorf("value do not implement protobuf Message interface")
	}
}

func (pb *Pb) Unmarshal(data []byte, v interface{}) error {
	message, ok := v.(proto.Message)
	if ok {
		return proto.Unmarshal(data, message)
	} else {
		return fmt.Errorf("value do not implement protobuf Message interface")
	}
}

func (pb *Pb) ToString(v interface{}) (string, error) {
	message, ok := v.(proto.Message)
	if ok {
		bytes, err := protojson.Marshal(message)
		return string(bytes), err
	} else {
		return "", fmt.Errorf("value do not implement protobuf Message interface")
	}
}
