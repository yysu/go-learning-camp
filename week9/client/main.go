package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	processgoim "gitee.com/abelli8306/geekbang-go8/week9/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接服务端出错,错误原因: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("与服务端连接建立成功...")
	fmt.Printf("请发送信息\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		_, _ = conn.Write([]byte(text))
		readResp(conn)
	}
}

func readResp(conn net.Conn) {
	rr := bufio.NewReader(conn)
	decode := processgoim.GoImD.WithStream(rr).Decode()
	result, err := decode.Result()
	if err != nil {
		fmt.Printf("err happen: %v", err)
	}
	fmt.Println(result.Pretty())
}
