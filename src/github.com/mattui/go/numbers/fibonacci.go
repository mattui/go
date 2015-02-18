package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Fibonacci(n uint) uint {
	result := []uint{0, 1}
	if n < 2 {
		return result[n]
	}

	var f0 uint = 0
	var f1 uint = 1
	var f2 uint = 0

	for i := uint(2); i <= n; i++ {
		f0 = f1 + f2
		f2 = f1
		f1 = f0
	}

	return f0
}

func main() {
	fmt.Println("Enter a number:")

	in := bufio.NewReader(os.Stdin)
	var num uint
	_, err := fmt.Fscanf(in, "%d\n", &num)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("f(%[1]d) = %[2]d\n", num, Fibonacci(num))
	}
}
