package client

import (
	"math/rand"
	"time"
)

func bytesGenerator(kb int) []byte {
	rand.Seed(time.Now().UnixNano())

	var list = make([]byte, 1024*kb)
	for i := range list {
		list[i] = byte(rand.Intn(256))
	}
	return list
}
