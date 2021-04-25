package main

import (
	"bufio"
	"fmt"
	"github.com/warsky007/demo/jsonrpc/common"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os/exec"
)

func call(args common.Args, client *rpc.Client) {
	reply := new(common.Reply)

	if err := client.Call("Test", args, reply); err != nil {
		fmt.Printf("call method test fail: %v\n", err)
		return
	}
	fmt.Println(reply)
}

func main() {
	cmd := exec.Command("/tmp/client.sh")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatalln(err.Error())
	}

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println("receive from client:", string(scanner.Bytes()))
		}
	}()

	myIO := &common.MyReadWriteCloser{
		ReadCloser:  out,
		WriteCloser: in,
	}
	client := jsonrpc.NewClient(myIO)

	args := common.Args{
		A: 1,
		B: 2,
	}
	call(args, client)

	args = common.Args{
		A: 2,
		B: 5,
	}
	call(args, client)

	cmd.Process.Kill()
	cmd.Wait()
	client.Close()
}
