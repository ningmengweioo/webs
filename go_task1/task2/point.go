package main

import "fmt"

func pointerDemo(a *int) {

	*a += 10
	fmt.Printf("函数内部：指针地址=%p，指针指向的值=%d\n", a, *a)
}
func unpointerDemo(a int) int {
	a += 10
	return a
}
func doubleEachElement(slicePtr *[]int) {
	slice := *slicePtr

	// 遍历切片，将每个元素乘以2
	for i := range slice {
		slice[i] *= 2
	}

}
func main() {
	// 	1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	// 考察点 ：指针的使用、值传递与引用传递的区别。
	t := 10
	pointerDemo(&t)
	fmt.Println(t)
	t1 := 10
	unpointerDemo(t1)
	fmt.Println(t1)

	// 2	题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	// 考察点 ：指针运算、切片操作。

	// 创建一个整数切片
	numbers := []int{2, 4, 6, 8}
	fmt.Printf("原始切片: %v\n", numbers)

	doubleEachElement(&numbers)

	// 打印修改后的切片
	fmt.Printf("元素乘以2后: %v\n", numbers)

}
