package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}

}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	//c := wg(chan int)
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w) //this is the real worker
	return w
}

//func chanDemo() {
//	var workers [10]worker
//	for i := 0; i < 10; i++ {
//		workers[i] = createWorker(i)
//	}
//	for i, worker := range workers {
//		worker.in <- 'a' + i
//
//	}
//	for _, worker := range workers {
//		<-worker.done
//	}
//
//	for i, worker := range workers {
//		worker.in <- 'A' + i
//
//	}
//	// wiat for all of them
//	for _, worker := range workers {
//		<-worker.done
//	}
//}
func ChanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i

	}
	for i, worker := range workers {
		worker.in <- 'A' + i

	}
	wg.Wait()

}
func main() {
	ChanDemo()
}
