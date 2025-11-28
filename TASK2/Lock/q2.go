package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func counter(idx int, a *int32) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(a, 1)
		fmt.Println("counter:", idx, *a)
	}
}

func main() {
	a := int32(0)
	for i := 0; i < 10; i++ {
		go counter(i, &a)
	}

	select {
	case <-time.After(time.Second):
		fmt.Println(a)
	}
}
