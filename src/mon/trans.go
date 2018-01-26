package m

import (
	"log"
	"os"
	"encoding/json"
	"fmt"
	"net"
)
const defaultBufSize = 4096


type FileInfo struct {
	Name string
	Size int
	Message []byte
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}


func upload_to_server(conn net.Conn, filepath string) bool {

	var file *os.File

	err := checkFileIsExist(filepath)
	if err != true {
		log.Fatal("file not exist")
	}
	file, er := os.Open(filepath)
	if er != nil {
		log.Println(filepath + "cant open")
	}

	var temp_size int = 0
	temp_file := make([]byte, defaultBufSize)

	for {
		n, _ := file.Read(temp_file)
		if 0 == n {
			break
		}
		temp_size += n
	}

	fi := &FileInfo {
		filepath,
		temp_size,
		temp_file,
	}

	b, er := json.Marshal(fi)
	if er != nil {
		log.Fatal("marshal")
	}
	fmt.Println(string(b))
	conn.Write(b)
	return true
}
