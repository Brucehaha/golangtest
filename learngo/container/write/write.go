package write

import (
	"bufio"
	"fmt"
	"os"

	fib "brucego.com/learngo/container/functional"
)

func WriteTestFile(filename string) {
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
	defer writer.Flush()
	fib := fib.Fibnonacchi()

	reader := bufio.NewScanner(fib)
	for reader.Scan() {
		fmt.Fprintln(writer, reader.Text())
	}
}
