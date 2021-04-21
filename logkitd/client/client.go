package main

import (
	"fmt"
	"github.com/warsky007/demo/logkitd/common/ipc"
	"github.com/warsky007/demo/logkitd/common/pb"
	"log"
	"os"
	"runtime"
)

func handle(req *pb.Request) (rsp *pb.Response, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = fmt.Errorf("handle ipc message %d panic: %v", req.GetType(), err)
			log.Printf("handle ipc message %d panic: %v", req.GetType(), err)
		}
	}()

	rsp = &pb.Response{
		Type: req.GetType(),
	}

	switch req.GetType() {
	case pb.TypeName_GetOs:
		rsp.Os = &pb.OsInfoRsp{
			Os:   runtime.GOOS,
			Arch: runtime.GOARCH,
		}
	case pb.TypeName_GetPid:
		pid := os.Getpid()
		rsp.Pid = &pb.PidRsp{
			Pid: int32(pid),
		}
	default:
		err = fmt.Errorf("got unknown ipc type %d", req.GetType())
	}

	return
}

func main() {
	pipe := os.Stdout
	os.Stdout = os.Stderr

	serialize := pb.NewPb()
	ipcs := ipc.CreateIpc(os.Stdin, pipe, serialize)
	ipcs.Start()

	for msg := range ipcs.RecvCh {
		req := &pb.Request{}
		err := ipcs.Unmarshal(msg.Data, req)
		if err != nil {
			log.Printf("can't decode msg: %v", err)
			continue
		}

		rsp, err := handle(req)
		if err != nil {
			rsp.Error = err.Error()
		}
		rspMsg, err := ipcs.ToMessage(rsp, msg.MsgId)
		if err != nil {
			log.Printf("marshal response msg fail: %v", err)
			continue
		}
		ipcs.SendMsg(rspMsg)
	}

	ipcs.Stop()
}
