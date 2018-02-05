package main

import (
	"net"
	"fmt"
	"log"
	"reflect"
	"io"
	"../mon"
	"runtime"
)



const server_addr = "127.0.0.1:9000"

func main()  {

	//启动本地监控程序
	go m.Client_monitor()

	//打印客户端信息
	var goos = runtime.GOOS
	fmt.Printf("The os is : %s\n", goos)

	//连接服务器
	conn, err := net.Dial("tcp", server_addr)
	if err != nil {
		log.Fatal("cilent dail")
	}

	//打印服务段地址、本机地址、通讯协议
	fmt.Println("Connect to server: " + conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.LocalAddr()))
	buf := make([]byte, 10)


	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}

		fmt.Print(string(buf[:n]))
	}
}


