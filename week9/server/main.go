package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/Terry-Mao/goim/api/protocol"
	"github.com/Terry-Mao/goim/pkg/bytes"
)

func main() {
	socket, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("开启监听失败,错误原因: ", err)
		return
	}
	defer socket.Close()
	fmt.Println("开启监听...")

	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("建立链接失败,错误原因: ", err)
			return
		}
		defer conn.Close()
		fmt.Println("建立链接成功,客户端地址是: ", conn.RemoteAddr())
		go logic(conn)

	}
}

func logic(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		slice, _ := reader.ReadSlice('\n') // 演示, 忽略异常
		cmd := strings.TrimSpace(string(slice))
		fmt.Println(cmd)
		var writer *bytes.Writer
		writer = reply(1, 4, 4, cmd)
		_, _ = conn.Write(writer.Buffer())
		time.Sleep(time.Second * 1)
	}
}

func reply(ver, op, seq int32, body string) *bytes.Writer {
	bs := []byte(body)
	writer := bytes.NewWriterSize(len(body) + 64)
	proto := &protocol.Proto{
		Ver:  ver,
		Op:   op,
		Seq:  seq,
		Body: bs,
	}
	proto.WriteTo(writer)
	return writer
}
