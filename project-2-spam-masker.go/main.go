package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Provide information to mask")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buffer  = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buffer = append(buffer, text[i])

		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
		}
	}

	fmt.Println(string(buffer))
}

