package main

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"

	//"encoding/binary"
	"fmt"
	pb "go-work/sixth_go_work/socket_msg"
	"io"
	"net"
)
const (
	TCP_HEADER_LEN = 9
	TCP_DATA_LEN = 4
)

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}


func main() {
	// 监听端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Listen Error: %s\n", err)
		return
	}
	// 监听循环
	for {
		// 接受客户端链接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Accept Error: %s\n", err)
			continue
		}
		// 处理客户端链接
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	// 关闭链接
	defer conn.Close()
	// 客户端
	fmt.Printf("Client: %s\n", conn.RemoteAddr())
	// 消息缓冲
	msgbuf := bytes.NewBuffer(make([]byte, 0, 10240))
	// 数据缓冲
	databuf := make([]byte, 4096)

	for {
		// 读取数据
		n, err := conn.Read(databuf)
		if err == io.EOF {
			fmt.Printf("Client exit: %s\n", conn.RemoteAddr())
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return
		}

		n, err = msgbuf.Write(databuf[:n])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return
		}
		// 消息分割循环
		for {
			////固定长度
			//if msgbuf.Len() <= 256 && msgbuf.Len() > 0{
			//	fmt.Printf("Client messge: %s\n", string(msgbuf.Next(n)))
			//}
			//
			////分隔符  msgbuf 消息只能读取一次
			//str := string(msgbuf.Next(n))
			//if strings.Contains(str,"??"){
			//	fmt.Printf("Client messge: %s\n", str[:(n-len("??"))])
			//}
			//
			//
			//if msgbuf.Len() == 0{
			//	break
			//}
			//
			////包头，包体
			//length := len(databuf)
			//fmt.Printf("databuf len:",len(databuf))
			//
			//var i int
			//for i = 0; i < length; i++ {
			//	if length < i + TCP_HEADER_LEN + TCP_DATA_LEN {
			//		break
			//	}
			//	if string(databuf[i:i + TCP_HEADER_LEN]) == "TCPHEADER" {
			//		msgLen := BytesToInt(databuf[i + TCP_HEADER_LEN : i + TCP_HEADER_LEN + TCP_DATA_LEN])
			//		if length < i + TCP_HEADER_LEN + TCP_DATA_LEN + msgLen {
			//			break
			//		}
			//
			//		i += TCP_HEADER_LEN + TCP_DATA_LEN + msgLen - 1
			//	}
			//}
			//
			//fmt.Printf("message buffer %v\n",string(databuf[TCP_HEADER_LEN:i]))
			//4.协议解析
			msg := &pb.SocketMsg{}
			err := proto.Unmarshal(msgbuf.Next(n),msg)
			if err != nil{
				fmt.Printf("unmarshal err :%v\n", err)
			}

			fmt.Printf("message buffer result:%v\n",string(msg.GetBody()))
			break



		}
	}
}


