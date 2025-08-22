package main

import "fmt"

func main() {
	// 给定一个字符串，判断该字符串是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
	strs1 := []string{"test", "te", "tes"}
	fmt.Printf("最长公共前缀: %s\n", longestCommonPrefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Printf("最长公共前缀: %s\n", longestCommonPrefix(strs2))
}

func longestCommonPrefix(strs []string) string {
	// 处理边界情况
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		// 比较当前字符串与基准字符串的每个字符
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		// 更新公共前缀
		prefix = prefix[:j]

		// 如果公共前缀为空，提前返回
		if prefix == "" {
			break
		}
	}

	return prefix
}
