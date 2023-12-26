package main

import (
	"flag"
	"mutliClientTest/internal/app/mct/client"
	"mutliClientTest/internal/app/mct/server"
)

func main() {
	serverPort := flag.Int("s", -1, "Порт сервера")
	clientAddr := flag.String("c", "", "Адрес сервера")
	clientCount := flag.Int("t", 50, "Количество клиентов")
	flag.Parse()

	if *serverPort != -1 {
		server.Run(serverPort)
	}

	if *clientAddr != "" {
		client.Run(*clientAddr, *clientCount)
	}
}
