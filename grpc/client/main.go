package main

import (
	"context"
	"github.com/warsky007/demo/grpc/common"
	"github.com/warsky007/demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"runtime"
	"sync"
)

type RpcServer struct {}

func(s *RpcServer)GetOsInfo(context.Context, *pb.Request) (*pb.OsInfoRsp, error){
	return &pb.OsInfoRsp{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}, nil
}

func(s *RpcServer)GetPid(context.Context, *pb.Request) (*pb.PidRsp, error){
	return &pb.PidRsp{
		Pid: int32(os.Getpid()),
	}, nil
}

func startRpcServer(listener net.Listener, wg sync.WaitGroup) {
	defer wg.Done()
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &RpcServer{})
	if err := s.Serve(listener); err != nil {
		log.Printf("serve fail: %v\n", err)
	}
}

func main() {
	pipe := os.Stdout
	os.Stdout = os.Stderr
	wg := sync.WaitGroup{}

	stdioConn := common.NewStdioConn(os.Stdin, pipe)
	listener := common.NewStdioListener()
	wg.Add(1)
	go startRpcServer(listener, wg)
	listener.Ready(stdioConn)
	wg.Wait()
}
