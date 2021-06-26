package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() { // race condition
			for {
				a[i]++
				runtime.Gosched()
			}
		}() //race condition here when i= 10,if not pass i in this function, as it is  anonymous function
	}
	time.Sleep(time.Millisecond)
	fmt.Print(a)
}
