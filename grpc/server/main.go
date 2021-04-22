package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/warsky007/demo/grpc/common"
	"github.com/warsky007/demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("/tmp/client")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Fatalf("sub process start fail: %v\n", err)
	}

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println("receive from client:", string(scanner.Bytes()))
		}
	}()

	stdioConn := common.NewStdioConn(out, in)
	conn, err := grpc.Dial("", grpc.WithInsecure(),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return stdioConn, nil
		}))
	if err != nil {
		log.Fatalf("connect fail: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	rsp, err := client.GetPid(ctx, &pb.Request{})
	if err != nil {
		fmt.Printf("get os info fail: %v", err)
	}
	fmt.Println(rsp.String())

	cmd.Wait()
}
