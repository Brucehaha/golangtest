package main

import (
	"fmt"

	fib "brucego.com/learngo/container/functional"
)

func main() {
	a := fib.Adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("%d-->%d ", i, s)

	}

	f := fib.Fibnonacchi()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fib.PrintFileContents(fib.Fibnonacchi())
}
