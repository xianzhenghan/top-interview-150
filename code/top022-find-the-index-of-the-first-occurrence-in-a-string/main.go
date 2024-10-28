package main

import (
	"fmt"
)

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
//
//
//
//示例 1：
//
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
//示例 2：
//
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
//
//提示：
//
//1 <= haystack.length, needle.length <= 104
//haystack 和 needle 仅由小写英文字符组成
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	next := make([]int, len(needle))
	for i := 0; i < len(next); i++ {
		next[i] = getMaxCommonLen(needle[:i+1])
	}
	fmt.Println(next)
	k := 0
	j := 0
	for j = 0; j < len(haystack); j++ {
		tmpj := j
		for ; k < len(needle) && tmpj < len(haystack); k++ {
			fmt.Println(k, string(needle[k]), tmpj, string(haystack[tmpj]))
			if needle[k] == haystack[tmpj] && k == len(needle)-1 {
				return tmpj - (len(needle) - 1)
			} else if needle[k] == haystack[tmpj] {
				tmpj++
			} else {
				k = 0
				break
			}
		}
		k = next[k]
	}

	return -1
}

// "前缀"和"后缀"的最长的共有元素的长度
// preIdx 范围 [0,len(needle)-1]
func getMaxCommonLen(needle string) int {
	postIdx := len(needle) - 1
	prdIdx := 0
	for i := len(needle) - 2; i >= 0; i-- {
		//fmt.Println(i, needle[:prdIdx+i+1], " <=> ", needle[postIdx-i:], len(needle[:prdIdx+i+1]))
		if needle[:prdIdx+i+1] == needle[postIdx-i:] {
			return len(needle[:prdIdx+i+1])
		}
	}
	return 0
}

func main() {
	//fmt.Println(getMaxCommonLen("ABCDAB"))
	fmt.Println(strStr("vbabba", "bbb"))
}
