package main

import (
	"fmt"
	"math"
)

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

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
