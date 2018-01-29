package main

import (
	"net"
	"log"
	"time"
	"fmt"
	"encoding/json"
)


type FileInfo struct {
	Name string
	Size int
	Message []byte
}
const serve_addr = "127.0.0.1:9000"
const defaultBufSize = 10240

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serve_addr)
	if err != nil {
		log.Fatal("resolveTcp")
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("listenTcp")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatal("accept")
		}
		go Handler_conn(conn)
	}
}

func Handler_conn(conn net.Conn) {

	//server func
	buf := make([]byte, defaultBufSize)
	n, _ := conn.Read(buf)

	up := &FileInfo{}
	err := json.Unmarshal([]byte(buf[:n]), &up)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(up.Message[:up.Size]))
	conn.Write([]byte("from server"))
	time.Sleep(time.Minute)
	conn.Close()
}