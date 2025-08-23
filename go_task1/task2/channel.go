package main

import (
	"fmt"
	"sync"
)

func Randnum(ch chan int) {
	defer close(ch)
	for i := 1; i < 11; i++ {
		fmt.Println(i)
		ch <- i
	}

}

func main() {
	// 	题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
	// 另一个协程从通道中接收这些整数并打印出来。
	// 考察点 ：通道的基本使用、协程间通信。
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)

	go func() {
		defer wg.Done()
		Randnum(ch)
	}()
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
	}()
	wg.Wait()

	// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 考察点 ：通道的缓冲机制。

	//分批用缓冲通道
	bufferedChannelDemo()

}

// 带缓冲通道的实现示例
func bufferedChannelDemo() {
	// 创建一个容量为10的带缓冲通道
	// 缓冲大小可以根据实际需求调整，这里选择10是因为它小于要发送的100个数据
	// 这样可以展示缓冲通道的工作原理
	bufferSize := 10
	ch := make(chan int, bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("生产者开始发送数据...")
		for i := 1; i <= 100; i++ {
			ch <- i
			// 每发送10个数据打印一次进度
			if i%10 == 0 {
				fmt.Printf("生产者已发送 %d 个数据，通道中当前数据量: %d\n", i, len(ch))
			}
		}
		close(ch) // 发送完成后关闭通道
		fmt.Println("生产者发送完毕并关闭通道")
	}()

	// 消费者协程：从通道接收并打印整数
	go func() {
		defer wg.Done()
		fmt.Println("消费者开始接收数据...")
		receivedCount := 0
		// 使用for range循环接收通道中的所有数据，直到通道关闭且为空
		for num := range ch {
			// 为了更清晰地看到缓冲效果，每接收5个数据打印一次
			if receivedCount%5 == 0 {
				fmt.Printf("消费者接收: %d, 通道中剩余数据量: %d\n", num, len(ch))
			}
			receivedCount++
		}
		fmt.Printf("消费者接收完毕，共接收 %d 个数据\n", receivedCount)
	}()

	wg.Wait() // 等待所有协程完成
	fmt.Println("带缓冲通道演示完成")
}
