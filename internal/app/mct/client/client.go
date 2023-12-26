package client

import (
	"fmt"
	"net"
	"sync"
)

func clientRoutine(wg *sync.WaitGroup, id int, addr string) {

	defer wg.Done()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := bytesGenerator(1)
	buffer := make([]byte, 1)
	for {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Client %d: Error sending message: %v\n", id, err)
			return
		}

		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Printf("Client %d: Error reading response: %v\n", id, err)
			return
		}
		Counter.Add(len(message))
	}
}

func Run(addr string, count int) {
	if count == 0 {
		count = 50
	}
	var wg sync.WaitGroup

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go clientRoutine(&wg, i, addr)
	}

	go Counter.TickCounter()

	wg.Wait()
}
