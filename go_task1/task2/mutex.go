package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func mutexCounterDemo1() {
	var counter int
	var mutex sync.Mutex
	var wg sync.WaitGroup
	const (
		goCount = 10
		incre   = 1000
	)
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("协程%d开始\n", id)
			for j := 0; j < incre; j++ {

				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			fmt.Printf("协程%d执行完毕\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("计数器的值为：%d\n", counter)

}
func mutexCounterDemo2() {
	var counter int32
	//var mutex sync.Mutex
	var wg sync.WaitGroup
	const (
		goCount = 10
		incre   = 1000
	)
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Printf("协程%d开始\n", id)
			for j := 0; j < incre; j++ {
				atomic.AddInt32(&counter, 1)
			}
			fmt.Printf("协程%d执行完毕\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Printf("计数器的值为：%d\n", counter)

}
func main() {
	// 1.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ： sync.Mutex 的使用、并发数据安全。
	mutexCounterDemo1()
	//2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ：原子操作、并发数据安全
	mutexCounterDemo2()

}
