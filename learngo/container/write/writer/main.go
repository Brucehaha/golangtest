package main

import (
	"bufio"
	"fmt"

	writer "brucego.com/learngo/container/write"

	fib "brucego.com/learngo/container/functional"
)

func main() {
	fib := fib.Fibnonacchi()
	scanner := bufio.NewScanner(fib)
	fmt.Println("yes")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	writer.WriteTestFile("abc.txt")
}
