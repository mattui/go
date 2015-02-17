package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Reverse(s string) string {
	r := []rune(s) // convert string to runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // swap runes
	}
	return string(r) // convert runes to string
}

func main() {
	fmt.Println("Enter a string:")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Reversed string:", Reverse(line))
	}
}
