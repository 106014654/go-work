package main

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"

	//"encoding/binary"
	"fmt"
	pb "go-work/sixth_go_work/socket_msg"
	"net"

	//"time"
)

var MaxNum = 256

const (
	TCP_HEADER = "TCPHEADER"
)

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)

	return bytesBuffer.Bytes()
}

func Enpack(msg []byte) []byte {
	return append(append([]byte(TCP_HEADER), IntToBytes(len(msg))...), msg...)
}

func main() {
	// 链接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Dial error: %s\n", err)
		return
	}
	// 客户端信息
	fmt.Printf("Client: %s\n", conn.LocalAddr())
	// 消息缓冲
	//msgbuf := bytes.NewBuffer(make([]byte, 0, 1024))
	msg := "我是utf-8的消息"



	//1.固定长度（每次发放固定长度：256)
	// 消息内容
	//message := []byte(msg)
	//// 消息长度
	//messageLen := len(message)
	//
	//totalMessage := make([]byte,0,MaxNum)
	//
	//
	//if messageLen <= MaxNum {
	//	for _,v := range message{
	//		totalMessage = append(totalMessage,v)
	//	}
	//
	//	for i:=0;i<MaxNum-messageLen;i++{
	//		totalMessage = append(totalMessage,byte(0))
	//	}
	//}

	//2.分隔符 ？？
	//msg += "??"
	//message := []byte(msg)
	//fmt.Printf("Client messge and length : %s\n" ,string(message))


	//3.设置包头，包体
	//message := []byte(msg)
	//result := Enpack(message)
	//fmt.Printf("Client messge and length : %s\n" ,string(result))

	//4.协议解码
	pbMessage := &pb.SocketMsg{
		Ver: 0,
		Op: 1,
		Seq: 2,
		Body: []byte(msg),
	}
fmt.Printf("pb message:%v",pbMessage)
	result, _ := proto.Marshal(pbMessage)

	conn.Write(result)
	conn.Close()
}
