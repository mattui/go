package main

import (
	"bufio"
	"fmt"
	"github.com/mattui/go/utils"
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
		count := 0

		for _, searchRune := range search {
			for _, runeValue := range vowels {
				if searchRune == runeValue {
					count++
				}
			}
		}

		fmt.Printf("%[1]d vowel%[2]s found\n", count, utils.Ternary(count == 1, "", "s"))
	}
}
