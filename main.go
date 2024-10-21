package main

import (
	"fmt"
	"math"
	"strings"
)

//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
//单词 是指仅由字母组成、不包含任何空格字符的最大
//子字符串
//。
//
//
//
//示例 1：
//
//输入：s = "Hello World"
//输出：5
//解释：最后一个单词是“World”，长度为 5。
//示例 2：
//
//输入：s = "   fly me   to   the moon  "
//输出：4
//解释：最后一个单词是“moon”，长度为 4。
//示例 3：
//
//输入：s = "luffy is still joyboy"
//输出：6
//解释：最后一个单词是长度为 6 的“joyboy”。
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅有英文字母和空格 ' ' 组成
//s 中至少存在一个单词

func lengthOfLastWord(s string) int {
	currentLen := 0
	s = strings.TrimRight(s, " ")
	fmt.Println(s)
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			return currentLen
		}
		currentLen++
	}
	return currentLen
}

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//提示：
//
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成
func longestCommonPrefix(strs []string) string {
	maxLen := math.MaxInt32
	// 所有len的最小值
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < maxLen {
			maxLen = len(strs[i])
		}
	}
	if maxLen == 0 {
		return ""
	}
	dp := make([]string, maxLen+1)
	dp[0] = ""
	for i := 0; i < maxLen; i++ {
		var j int
		for j = 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				return dp[i]
			}
		}

		if j == len(strs) {
			dp[i+1] = dp[i] + string(strs[0][i])
		}
		//fmt.Println(i, j, dp)
	}
	return dp[maxLen]
}

//2.o(n) 复杂度计算最长和为0的子串
//例如 [-2,2,3,-3,4]的最长子串是[-2,2,3,-3], 长度是4
func longestZero(nums []int) int {
	length := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == 0 && i+1 > length {
			length = i + 1
		}
	}
	return length
}

// word hello -> hello word
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	return reverseStr(s)
}

func reverseStr(s string) string {
	if len(s) == 0 {
		return ""
	}
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			idx = i
		}
	}
	return reverseWords(s[idx:]) + s[:idx]
}

func main() {
	//strs := []string{"cir", "car"}
	fmt.Println(reverseWords("hello word"))
}
