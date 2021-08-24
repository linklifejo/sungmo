// TCP Server
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"log"
	"net"
	"sync"
)

func init() {
	gob.Register(MyMsgBodyPing{})
}

func main() {
	l, err := net.Listen("tcp", ":5032")
	if nil != err {
		log.Fatalf("fail to bind address to 5032; err: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}
		go ConnHandler(conn)
	}
}

type MyMsg struct {
	Header MyMsgHeader
	Body   interface{}
}
type MyMsgHeader struct {
	MsgType1  string
	MsgType2  string
	MsgType3  string
	MsgType4  string
	MsgType5  string
	MsgType6  string
	MsgType7  string
	MsgType8  string
	MsgType9  string
	MsgType10 string
	MsgType11 string
	MsgType12 string

	Date string
}

type MyMsgBodyPing struct {
	Content string
}

type MyConnection struct {
	recv            chan MyMsg
	conn            net.Conn
	dec             *gob.Decoder
	codecBuffer     bytes.Buffer
	codecBufferLock sync.Mutex
	recvBuf         []byte
}

func (mc *MyConnection) RecvMessage() (MyMsg, error) {
	mc.codecBufferLock.Lock()
	defer mc.codecBufferLock.Unlock()

	lengthBuf := make([]byte, 4)
	_, err := mc.conn.Read(lengthBuf)
	if nil != err {
		return MyMsg{}, err
	}
	msgLength := binary.LittleEndian.Uint32(lengthBuf)

	mc.codecBuffer.Reset()

	for 0 < msgLength {
		n, err := mc.conn.Read(mc.recvBuf)
		if nil != err {
			return MyMsg{}, err
		}
		if 0 < n {
			data := mc.recvBuf[:n]
			mc.codecBuffer.Write(data)
			msgLength -= uint32(n)
		}
	}

	msg := MyMsg{}
	if err = mc.dec.Decode(&msg); nil != err {
		log.Panic("failed to decode message; err: %v", err)
		return msg, err
	}
	return msg, nil
}

func (mc *MyConnection) MyMsgReceiver() {
	// receiver
	go func() {
		for {
			msg, err := mc.RecvMessage()
			if nil != err {
				if io.EOF == err {
					log.Printf("connection is closed from client; %v", mc.conn.RemoteAddr().String())
					return
				}
				log.Printf("failed to recv message! err: %v", err)
				continue
			}

			mc.recv <- msg
		}
	}()

	// message handler
	for {
		select {
		case msg := <-mc.recv:
			log.Println(msg)
		}
	}
}

func ConnHandler(conn net.Conn) {
	mc := MyConnection{
		conn:    conn,
		recvBuf: make([]byte, 4096),
		recv:    make(chan MyMsg),
	}
	mc.dec = gob.NewDecoder(&mc.codecBuffer)

	mc.MyMsgReceiver()
}
