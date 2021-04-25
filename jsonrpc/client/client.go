package main

import (
	"fmt"
	"github.com/warsky007/demo/jsonrpc/common"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Server struct{}

func (s *Server) Test(args common.Args, reply *common.Reply) error {
	reply.Message = "hehe"
	reply.Data = make(map[string]interface{})
	reply.Data["value"] = args.A + args.B
	return fmt.Errorf("this is a test error")
}

func main() {
	pipe := os.Stdout
	os.Stdout = os.Stderr

	myIO := &common.MyReadWriteCloser{
		ReadCloser:  os.Stdin,
		WriteCloser: pipe,
	}
	server := rpc.NewServer()
	codec := jsonrpc.NewServerCodec(myIO)
	server.Register(&Server{})
	server.ServeCodec(codec)
}
