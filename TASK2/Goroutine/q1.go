package main

import (
	"fmt"
	"time"
)

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/

func printEven(ch chan<- int) {
	for i := 1; i < 11; i += 2 {
		fmt.Println(i)
	}
	ch <- 2
}

func printOdd(ch chan<- int) {
	for i := 2; i < 11; i += 2 {
		fmt.Println(i)
	}
	ch <- 1
}

func main() {
	ch := make(chan int)
	go printEven(ch)
	go printOdd(ch)

	doneJobCount := 0
	timeout := time.After(5 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("recieve one done signal:", v)
				doneJobCount++
			}
		case <-timeout:
			fmt.Println("Timeout")
			return
		}
		if doneJobCount >= 2 {
			fmt.Println("All job done")
			break
		}
	}
}
