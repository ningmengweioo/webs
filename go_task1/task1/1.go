package main

import "fmt"

func main() {
	// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素
	check := []int{5, 5, 6, 6, 7, 7, 8}

	//single := make(map[int]int)
	single := map[int]int{}
	for _, v := range check {
		single[v]++
	}
	fmt.Println(single)
	for k, v := range single {
		if v == 1 {
			fmt.Printf("single number is : %d \n", k)
		}
	}
	fmt.Println("----------")
}
