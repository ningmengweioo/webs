package main

import "fmt"

func main() {
	digits1 := []int{1, 2, 3}
	fmt.Printf("%v 加一的结果: %v\n", digits1, plusOne(digits1))

	digits2 := []int{1, 2, 9}
	fmt.Printf("%v 加一的结果: %v\n", digits2, plusOne(digits2))

	digits3 := []int{9, 9, 9}
	fmt.Printf("%v 加一的结果: %v\n", digits3, plusOne(digits3))

	digits4 := []int{4, 3, 2, 1}
	fmt.Printf("%v 加一的结果: %v\n", digits4, plusOne(digits4))

	// 额外测试用例：验证全0数组
	digits5 := []int{0, 0, 0}
	fmt.Printf("%v 加一的结果: %v\n", digits5, plusOne(digits5))
}

func plusOne(digits []int) []int {
	// 创建数组副本，避免修改原始数组
	result := make([]int, len(digits))
	//todo 问题的重点
	copy(result, digits)

	n := len(result)

	// 从最后一位开始处理
	for i := n - 1; i >= 0; i-- {
		// 如果当前位加1后不等于10，不需要进位，直接返回结果
		result[i]++
		result[i] %= 10

		// 如果当前位不是0，说明没有进位，直接返回
		if result[i] != 0 {
			return result
		}
		// 否则继续向前处理进位
	}

	// 如果所有位都是9，需要在数组前面添加一个1
	// 例如 [9,9,9] -> [1,0,0,0]
	newResult := make([]int, n+1)
	newResult[0] = 1
	// 其余位已经是0，不需要额外设置

	return newResult
}
