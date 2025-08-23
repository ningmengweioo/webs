package main

import (
	"fmt"
	"sync"
	"time"
)

func write1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Printf("i is : %d \n", i)
		}
	}
}

func write2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("i is : %d \n", i)
		}
	}
}

type Task func()

type TaskResult struct {
	TaskIndex     int
	ExecutionTime time.Duration
}

func TaskScheduler(tasks []Task) []TaskResult {
	results := make([]TaskResult, len(tasks))
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func(index int, t Task) {
			defer wg.Done()
			startTime := time.Now()
			t()
			results[index] = TaskResult{
				TaskIndex:     index,
				ExecutionTime: time.Since(startTime),
			}
		}(i, task)

	}
	wg.Wait()
	return results
}

func main() {
	//1. 协程的使用
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go write1(&wg)
	// go write2(&wg)
	// wg.Wait()

	//2.协程的调度器
	tasks := []Task{
		// 任务1：模拟耗时操作
		func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("任务1执行完成")
		},
		// 任务2：模拟耗时操作
		func() {
			time.Sleep(150 * time.Millisecond)
			fmt.Println("任务2执行完成")
		},
		// 任务3：模拟耗时操作
		func() {
			time.Sleep(80 * time.Millisecond)
			fmt.Println("任务3执行完成")
		},
	}

	// 执行任务调度器
	fmt.Println("开始执行任务...")
	results := TaskScheduler(tasks)
	fmt.Println("所有任务执行完成，结果如下：")

	for _, result := range results {
		fmt.Printf("任务 %d 执行时间: %v\n", result.TaskIndex, result.ExecutionTime)
	}
}
