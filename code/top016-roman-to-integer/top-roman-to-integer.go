package main

import "fmt"

//https://leetcode.cn/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150

//提示：
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

//1 <= s.length <= 15
//s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
//题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
//题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
//IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
//关于罗马数字的详尽书写规则，可以参考 罗马数字 - 百度百科。
func romanToInt(s string) int {
	transMap := make(map[string]int)
	// 单个字
	transMap["I"] = 1
	transMap["V"] = 5
	transMap["X"] = 10
	transMap["L"] = 50
	transMap["C"] = 100
	transMap["D"] = 500
	transMap["M"] = 1000
	//两个
	transMap["IV"] = 4
	transMap["IX"] = 9
	transMap["XL"] = 40
	transMap["XC"] = 90
	transMap["CD"] = 400
	transMap["CM"] = 900
	res := 0
	length := len(s)
	for idx := 0; idx < length; idx++ {
		if idx < length-1 {
			v, ok := transMap[s[idx:idx+2]]
			if ok {
				res += v
				idx++
				continue
			}
		}
		res += transMap[string(s[idx])]
	}
	return res
}

func main() {
	fmt.Println(romanToInt("MDCCCXCIX"))
}
