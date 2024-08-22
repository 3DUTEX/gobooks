package main

import (
	"fmt"
	"time"
)

func counter(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second) // await one second
	}
}

func worker(workerID int, data chan int) {
	for x:= range data {
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

// goroutine 1
func main() {
	ch := make(chan int)

	go worker(1, ch)
	go worker(2, ch)

	for i := range 10 {
		ch <- i
	}
}

// for create one goroutine, use a reserved word "go". example: "go counter(5)"
// for create one channel, use a reserverd word "make". example: channel := make(chan string)