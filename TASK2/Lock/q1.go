package main

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

import (
	"fmt"
	"sync"
	"time"
)

func counter(idx int, a *int, lock *sync.Mutex) {
	(*lock).Lock()
	defer (*lock).Unlock()
	for i := 0; i < 1000; i++ {
		*a++
		fmt.Println("counter:", idx, *a)
	}
}

func main() {
	a := 0
	lock := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go counter(i, &a, &lock)
	}
	select {
	case <-time.After(5 * time.Second):
	}
	fmt.Println(a)
}
