package client

func bytesGenerator(kb int) []byte {
	var list = []byte{}
	for i := 0; i < 1024*kb; i++ {
		list = append(list, 0)
	}
	return list
}
