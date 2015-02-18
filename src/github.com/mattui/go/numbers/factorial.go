package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return Factorial(n-1) * n
}

func main() {
	fmt.Println("Enter a number:")

	in := bufio.NewReader(os.Stdin)
	var num int
	_, err := fmt.Fscanf(in, "%d\n", &num)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%[1]d! = %[2]d\n", num, Factorial(num))
	}
}
