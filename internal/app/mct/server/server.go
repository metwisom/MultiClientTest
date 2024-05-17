package server

import (
	"fmt"
	"mutliClientTest/internal/pkg/printer"
	"net"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
	}
}

func Run(port ...*int) {
	var targetPort = ":8080"
	if len(port) == 1 && port[0] != nil {
		targetPort = ":" + strconv.Itoa(*port[0])
	}
	listener, err := net.Listen("tcp", targetPort)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	printer.ListenMessage(targetPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
