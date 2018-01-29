package main

import (
"fmt"
"github.com/smallfish/ftp"
"os"
)

func main() {
	ftp := new(ftp.FTP)
	// debug default false
	ftp.Debug = true
	ftp.Connect("47.94.137.217", 21)

	// login
	ftp.Login("Sequin_YF", "ting199787*/")
	if ftp.Code == 530 {
		fmt.Println("error: login failure")
		os.Exit(-1)
	}

	// pwd
	ftp.Pwd()
	fmt.Println("code:", ftp.Code, ", message:", ftp.Message)

	// make dir
	//ftp.Mkd("~/")
	//ftp.Request("TYPE I")
	fmt.Println("list")
	ftp.List()
	fmt.Println("listover")
	// stor file
	//b, _ := ioutil.ReadFile("/Users/sequin_yf/go/src/CloudHybrid/sa")
	//ftp.Stor("/Users/sequin_yf/go/src/CloudHybrid/sa", b)
	ftp.Quit()
}