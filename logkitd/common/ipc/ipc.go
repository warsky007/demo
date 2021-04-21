package ipc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

func CreateIpc(read io.ReadCloser, write io.WriteCloser, serialize Serialization) *Ipc {
	return &Ipc{
		reader:    read,
		writer:    write,
		RecvCh:    make(chan *Message, 1),
		SendCh:    make(chan *Message, 1),
		status:    Running,
		wg:        sync.WaitGroup{},
		split:     SplitMessages,
		serialize: serialize,
	}
}

func (i *Ipc) Start() {
	i.wg.Add(1)
	go i.read()
	i.wg.Add(1)
	go i.write()
}

func SafeSend(ch chan *Message, value *Message) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}

func (i *Ipc) read() {
	defer i.wg.Done()

	scanner := bufio.NewScanner(i.reader)
	scanner.Split(i.split)

	for i.status == Running && scanner.Scan() {
		data := scanner.Bytes()
		msg, err := Decode(data)
		if err == nil {
			ch := i.RecvCh
			if tmp, ok := i.cache.Load(msg.MsgId); ok {
				ch = tmp.(chan *Message)
			}
			if closed := SafeSend(ch, msg); closed {
				log.Printf("[IPC] module's receive channel has been closed\n")
				return
			}
		} else {
			log.Printf("[IPC] module got invalid message: %v\n", err)
		}
	}

	log.Printf("[IPC] module's read gorountine exit\n")
}

func (i *Ipc) write() {
	defer i.wg.Done()

	for i.status == Running {
		msg, ok := <-i.SendCh
		if ok {
			bytes := Encode(msg)
			_, err := i.writer.Write(bytes)
			if err != nil {
				log.Printf("[IPC] module write message fail: %v\n", err)
			}
		} else {
			log.Printf("[IPC] module's send channel has been closed\n")
			log.Printf("[IPC] module's write gorountine exit\n")
			return
		}
	}

	log.Printf("[IPC] module's write gorountine exit\n")
}

func (i *Ipc) GetMsgId() uint32 {
	i.m.Lock()
	defer i.m.Unlock()
	i.id += 1
	return i.id
}

func (i *Ipc) Stop() {
	i.status = Closing
	close(i.SendCh)
	i.wg.Wait()
	close(i.RecvCh)
	// TODO close reader and writer
	// 因为项目中默认使用的是匿名管道，cmd模块会自动帮忙close，所以这里不需要close
	// 如果换用了其它的ipc方式，记得要close，否则会导致句柄泄露
	i.status = Closed
	return
}

func (i *Ipc) Marshal(v interface{}) ([]byte, error) {
	return i.serialize.Marshal(v)
}

func (i *Ipc) Unmarshal(data []byte, v interface{}) error {
	return i.serialize.Unmarshal(data, v)
}

func (i *Ipc) ToString(v interface{}) string {
	str, _ := i.serialize.ToString(v)
	return str
}

func (i *Ipc) ToMessage(v interface{}, id uint32) (*Message, error) {
	data, err := i.serialize.Marshal(v)
	if err != nil {
		return nil, err
	}

	if id == 0 {
		id = i.GetMsgId()
	}

	msg := &Message{
		MsgId: id,
		Data: data,
	}

	return msg, nil
}

func (i *Ipc) SendMsg(msg *Message) error {
	if closed := SafeSend(i.SendCh, msg); closed {
		return fmt.Errorf("[IPC] module's send channel has been closed, can't send")
	}

	return nil
}

func (i *Ipc) SendAndRecvBlock(msg *Message, timeout time.Duration) (*Message, error) {
	if err := i.SendMsg(msg); err != nil {
		return nil, err
	}

	rspCh := make(chan *Message)
	i.cache.Store(msg.MsgId, rspCh)

	select {
	case rsp := <-rspCh:
		i.cache.Delete(msg.MsgId)
		close(rspCh)
		return rsp, nil
	case <-time.After(timeout):
		i.cache.Delete(msg.MsgId)
		close(rspCh)
		err := fmt.Errorf("get response timeout")
		return nil, err
	}
}
