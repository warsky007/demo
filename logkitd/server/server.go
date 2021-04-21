package main

import (
	"bufio"
	"fmt"
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

	serialize := pb.NewPb()
	ipcs := ipc.CreateIpc(out, in, serialize)
	ipcs.Start()

	// block example
	osReq := &pb.Request{
		Type: pb.TypeName_GetOs,
	}
	msg, err := ipcs.ToMessage(osReq, 0)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := ipcs.SendAndRecvBlock(msg, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	rsp := &pb.Response{}
	err = ipcs.Unmarshal(resp.Data, rsp)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response error: %s", rsp.GetError())
	log.Printf("got pid %d\n", rsp.GetPid().GetPid())
	log.Printf("got os %s arch %s\n", rsp.GetOs().GetOs(), rsp.GetOs().GetArch())
	log.Println("message:", ipcs.ToString(rsp))

	cmd.Process.Kill()
	cmd.Wait()
	ipcs.Stop()
}
