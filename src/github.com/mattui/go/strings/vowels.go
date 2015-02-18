package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const vowels = "aeiou"

func main() {
	fmt.Println("Enter a string:")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		search := strings.ToLower(line)
		for _, runeValue := range vowels {
			strValue := string(runeValue)
			fmt.Printf("%[1]s = %[2]d\n", strValue, strings.Count(search, strValue))
		}
	}
}
