package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001") //tcp client
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close() //close TCP connection
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n') // read user inputs
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // input "q" then quit
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
	}
}
