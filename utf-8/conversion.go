package main

import "fmt"

func main() {
	inputWords := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var byteWords [][]byte
	for _, word := range inputWords {
		byteWord := []byte(word)
		fmt.Println(byteWord)
		byteWords = append(byteWords, byteWord)
	}

	for _, word := range byteWords {
		fmt.Println(string(word))
	}
}
