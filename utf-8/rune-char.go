package main

import "fmt"

func main() {
	const targetWord = "console"

	for _, char := range targetWord {
		fmt.Printf("%c\n", char)
		fmt.Printf("\tdecimal: %[1]d\n", char)
		fmt.Printf("\thex    : 0x%[1]x\n", char)
		fmt.Printf("\tbinary : 0b%08[1]b\n", char)
	}

	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
