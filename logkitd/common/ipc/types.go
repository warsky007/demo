package ipc

import (
	"bufio"
	"io"
	"sync"
)

type Status int

const (
	Running Status = iota
	Closing
	Closed
)

type Ipc struct {
	RecvCh chan *Message
	SendCh chan *Message
	reader io.ReadCloser
	writer io.WriteCloser
	status Status
	wg     sync.WaitGroup
	split  bufio.SplitFunc
	id     uint32
	m      sync.Mutex
	cache  sync.Map
}
