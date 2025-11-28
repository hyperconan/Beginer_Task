package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

func job(id int, ch chan<- map[int]time.Duration) {
	startTime := time.Now()
	defer func() {
		ch <- map[int]time.Duration{
			id: time.Since(startTime),
		}
	}()
	//job content
	sleepTime := rand.Intn(500)
	fmt.Println("id:", id, " sleep:", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	//job content

}

func main() {
	ch := make(chan map[int]time.Duration)
	jobSize := 10
	for jobId := 1; jobId <= jobSize; jobId++ {
		go job(jobId, ch)
	}

	doneJobCount := 0
	jobTimeCounter := make(map[int]time.Duration, jobSize)
	timeout := time.After(5 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if ok {
				doneJobCount++
				for jobid, duration := range v {
					jobTimeCounter[jobid] = duration
				}
			}
		case <-timeout:
			fmt.Println("Timeout")
			break
		}

		if doneJobCount >= jobSize {
			break
		}
	}

	for id, duration := range jobTimeCounter {
		fmt.Printf("JobID:%d Duration:%v \n", id, duration)
	}
}
