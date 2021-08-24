package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"net"
	"sync"
	"time"
)

func init() {
	gob.Register(MyMsgBodyPing{})
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
	send chan MyMsg

	conn            net.Conn
	enc             *gob.Encoder
	codecBuffer     bytes.Buffer
	codecBufferLock sync.Mutex
}

func (mc *MyConnection) SendMessage(msg MyMsg) error {
	mc.codecBufferLock.Lock()
	defer mc.codecBufferLock.Unlock()

	mc.codecBuffer.Reset()

	mc.enc.Encode(msg)

	lengthBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuf, uint32(mc.codecBuffer.Len()))

	if _, err := mc.conn.Write(lengthBuf); nil != err {
		log.Printf("failed to send msg length; err: %v", err)
		return err
	}

	if _, err := mc.conn.Write(mc.codecBuffer.Bytes()); nil != err {
		log.Printf("failed to send msg; err: %v", err)
		return err
	}

	return nil
}

func (mc *MyConnection) MyMsgSender() {
	for {
		select {
		case msg := <-mc.send:
			if err := mc.SendMessage(msg); nil != err {
				log.Printf("failed to send message; err: %v", err)
			}
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", ":5032")
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	mc := MyConnection{
		conn: conn,
		send: make(chan MyMsg),
	}
	mc.enc = gob.NewEncoder(&mc.codecBuffer)

	go mc.MyMsgSender()

	for {
		mc.send <- MyMsg{
			Header: MyMsgHeader{
				MsgType1:  "ping10000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType2:  "ping20000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType3:  "ping30000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType4:  "ping40000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType5:  "ping50000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType6:  "ping60000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType7:  "ping70000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType8:  "ping80000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType9:  "ping90000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType10: "ping10000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType11: "ping110000000000000000000000000000000000000000000000000000000000000000000000000",
				MsgType12: "ping120000000000000000000000000000000000000000000000000000000000000000000000000",

				Date: time.Now().UTC().Format(time.RFC3339),
			},
			Body: MyMsgBodyPing{
				Content: "Hello! I'm alive!00000000000000000000000000000000000000000000000000000000000000",
			},
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}
