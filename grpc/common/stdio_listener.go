package common

import (
	"net"
)

type StdioListener struct {
	closed   bool
	wait     chan struct{}
	onlyConn net.Conn
}

func NewStdioListener() *StdioListener {
	return &StdioListener{
		wait: make(chan struct{}),
	}
}

func (lis *StdioListener) Ready(conn net.Conn) {
	lis.onlyConn = conn
	lis.wait <- struct{}{}
}

func (lis *StdioListener) Accept() (net.Conn, error) {
	<- lis.wait
	return lis.onlyConn, nil
}

func (lis *StdioListener) Close() error {
	lis.closed = true
	close(lis.wait)
	return nil
}

func (lis *StdioListener) Addr() net.Addr {
	return NewStdioAddr("listener")
}