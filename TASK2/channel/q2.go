package main

import (
	"fmt"
	"time"
)

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func producer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Println("Recv:", i)
	}
}

func main() {
	ch := make(chan int, 4)
	go producer(ch)
	go consumer(ch)

	select {
	case <-time.After(time.Second * 6):
		fmt.Println("timeout")
	}
}
