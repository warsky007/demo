package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/warsky007/demo/logkitd/common/ipc"
	"github.com/warsky007/demo/logkitd/common/pb"
	"log"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("/tmp/client")
	in, _ := cmd.StdinPipe()
	out, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	time.Sleep(time.Second)

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println("receive from client:", string(scanner.Bytes()))
		}
	}()

	ipcs := ipc.CreateIpc(out, in, ipc.SplitMessages)
	err := ipcs.Start()
	if err != nil {
		log.Fatalf("start ipc module fail: %v", err)
	}

	// block example
	osReq := &pb.Request{
		Type: pb.TypeName_GetPid,
	}
	data, err := proto.Marshal(osReq)
	if err != nil {
		log.Fatal(err)
	}
	msg := &ipc.Message{
		Data: data,
	}

	resp, err := ipcs.SendAndRecvBlock(msg, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	rsp := &pb.Response{}
	err = proto.Unmarshal(resp.Data, rsp)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response error: %s", rsp.GetError())
	log.Printf("got pid %d\n", rsp.GetPid().GetPid())
	log.Printf("got os %s arch %s\n", rsp.GetOs().GetOs(), rsp.GetOs().GetArch())

	cmd.Process.Kill()
	cmd.Wait()
	ipcs.Stop()
}
