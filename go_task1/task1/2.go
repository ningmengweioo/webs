package main

import "fmt"

func main() {
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	testCases := []string{
		"()",     // true
		"()[]{}", // true
		"(]",     // false
		"([)]",   // false
		"{[]}",   // true
	}

	for _, tc := range testCases {
		result := isValid(tc)
		fmt.Printf("字符串 \"%s\" 是否有效: %t\n", tc, result)
	}

}
func isValid(s string) bool {

	// 创建一个栈来存储开括号
	stack := []rune{}

	// 创建一个映射，用于匹配闭括号和开括号
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是闭括号
		if openBracket, ok := pairs[char]; ok {
			// 检查栈是否为空或栈顶元素是否与当前闭括号匹配
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是开括号，入栈
			stack = append(stack, char)
		}
	}

	// 如果栈为空，说明所有括号都匹配成功
	return len(stack) == 0
}
