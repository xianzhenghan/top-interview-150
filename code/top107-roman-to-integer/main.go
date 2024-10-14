package main

import (
	"fmt"
)

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

//1 <= num <= 3999
func intToRoman(num int) string {
	transMap := make(map[int]string)
	// 单个字
	transMap[1] = "I"
	transMap[5] = "V"
	transMap[10] = "X"
	transMap[50] = "L"
	transMap[100] = "C"
	transMap[500] = "D"
	transMap[1000] = "M"
	transMap[4] = "IV"
	transMap[9] = "IX"
	transMap[40] = "XL"
	transMap[90] = "XC"
	transMap[400] = "CD"
	transMap[900] = "CM"
	v, ok := transMap[num]
	if ok {
		return v
	}
	ans := ""
	sacle := 1
	for num > 0 {
		yushu := (num % 10)
		fmt.Println(yushu * sacle)
		v, ok := transMap[yushu*sacle]
		if ok {
			ans = v + ans
		} else {
			// 2 3 ; 6 7 8 都需要在右侧加上字符串
			plus := transMap[sacle]
			mode := ""
			nums := 0
			if yushu < 4 { // (n个plus)
				nums = yushu - 1
				mode = transMap[sacle]
			} else { // 5*scale 右边加上  (n个plus)
				nums = yushu - 5
				mode = transMap[sacle*5]
			}
			for nums > 0 {
				mode = mode + plus
				nums--
			}
			if yushu > 0 {
				ans = mode + ans
			}

		}
		sacle = sacle * 10
		num = num / 10
	}
	return ans
}

//"MMMDCCXLIX"
func main() {
	fmt.Println(intToRoman(20))
}
