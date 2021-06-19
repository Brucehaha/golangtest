package main

import (
	"bufio"
	"fmt"
	"os"

	fib "brucego.com/learngo/container/functional"
)

func writerFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)

		} else {
			fmt.Printf("%s %s %s", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	fib := fib.Fibnonacchi()
	reader := bufio.NewScanner(fib)
	for i := 0; i < 20; i++ {
		fmt.Printf(reader.Text())
	}
}
func main() {
	fib := fib.Fibnonacchi()
	reader := bufio.NewScanner(fib)
	for i := 0; i < 20; i++ {
		fmt.Println(reader.Text())
	}
}
