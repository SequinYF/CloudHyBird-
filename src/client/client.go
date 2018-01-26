package main

import (
	"net"
	"fmt"
	"log"
	"reflect"
	"io"
	"../mon"
)



const server_addr = "127.0.0.1:9000"

func main()  {

	//go m.client_monitor()

	conn, err := net.Dial("tcp", server_addr)
	if err != nil {
		log.Fatal("cilent dail")
	}
	//

	fmt.Println("connect to server: " + conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr())
	fmt.Println(reflect.TypeOf(conn.LocalAddr()))
	//	conn.Write([]byte("from client"))
	buf := make([]byte, 10)
	//m.upload_to_server(conn, filepath)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
		}

		fmt.Print(string(buf[:n]))
	}
}


