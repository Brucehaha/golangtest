package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	g := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			g <- i
			i++
		}
	}()
	return g
}
func doWorker(id int, w chan int) {
	for n := range w {
		time.Sleep(1 * time.Second)

		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	w := make(chan int)
	go doWorker(id, w)
	return w
}
func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	n := 0
	tf := time.After(10 * time.Second)
	tt := time.Tick(time.Millisecond * 300)
	var values []int
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = worker

		}

		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-tf:
			fmt.Println("bye")
			return
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tt:
			fmt.Println(len(values))

		}

	}

}
