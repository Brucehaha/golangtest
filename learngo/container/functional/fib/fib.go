package main

import (
	"fmt"

	fib "brucego.com/learngo/container/functional"
)

func main() {
	f := fib.Fibnonacchi()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fib.PrintFileContents(fib.Fibnonacchi())
}
