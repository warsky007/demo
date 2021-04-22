package common

import (
	"io"
	"net"
	"time"
)

type StdioAddr struct {
	s string
}

func NewStdioAddr(s string) *StdioAddr {
	return &StdioAddr{s}
}
func (a *StdioAddr) Network() string {
	return "stdio"
}

func (a *StdioAddr) String() string {
	return a.s
}

type StdioConn struct {
	in     io.Reader
	out    io.Writer
	closed bool
	local  *StdioAddr
	remote *StdioAddr
}

func NewStdioConn(in io.Reader, out io.Writer) *StdioConn {
	return &StdioConn{
		local:  NewStdioAddr("local"),
		remote: NewStdioAddr("remote"),
		in:     in,
		out:    out,
	}
}

func (s *StdioConn) LocalAddr() net.Addr {
	return s.local
}

func (s *StdioConn) RemoteAddr() net.Addr {
	return s.remote
}

func (s *StdioConn) Read(b []byte) (n int, err error) {
	return s.in.Read(b)
}

func (s *StdioConn) Write(b []byte) (n int, err error) {
	return s.out.Write(b)
}

func (s *StdioConn) Close() error {
	s.closed = true
	return nil
}

func (s *StdioConn) SetDeadline(t time.Time) error {
	return nil
}

func (s *StdioConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (s *StdioConn) SetWriteDeadline(t time.Time) error {
	return nil
}
