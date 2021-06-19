package main

import (
	"bufio"
	"fmt"
	"os"

	fib "brucego.com/learngo/container/functional"
)

func writerFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)

		} else {
			fmt.Printf("%s %s %s", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Println(writer.Size())
	f := fib.Fibnonacchi()
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		//writer.WriteString(fmt.Sprint(f()))
		fmt.Fprintln(writer, f())
		//
	}
	fmt.Println(writer.Size())
}

func main() {
	writerFile("abc.txt")
}
