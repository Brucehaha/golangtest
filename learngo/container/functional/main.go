package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func Fibnonacchi() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type iAdder func(int) (int, iAdder)

func Adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, Adder2(base + v)
	}
}

func fibnonacchi2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	//toDO incorrect if p is too small
	return strings.NewReader(s).Read(p)
}
func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
