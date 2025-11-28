package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func consomer(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consomer(ch)

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
		return
	}
}
