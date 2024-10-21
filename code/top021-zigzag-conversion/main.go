package main

import (
	"fmt"
	"strings"
)

//1 <= s.length <= 1000
//s 由英文字母（小写和大写）、',' 和 '.' 组成
//1 <= numRows <= 1000
//https://leetcode.cn/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	vn := 2*numRows - 2
	res := make([][]string, numRows)
	//for i := 0; i < numRows; i++ {
	//	res[i] = make([]byte, 0)
	//}

	for i := 0; i < len(s); i++ {
		vnY := i % vn
		if vnY < numRows {
			res[vnY] = append(res[vnY], string(s[i]))
		} else {
			res[vn-vnY] = append(res[vn-vnY], string(s[i]))
		}
		fmt.Println(res)
	}
	str := ""
	for i := 0; i < len(res); i++ {
		str += strings.Join(res[i], "")
	}
	return str
}

func main() {
	fmt.Println(convert("ABCD", 3))
}
