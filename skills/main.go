package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	stra := "qwertyuiop"
	strb := "1234567890"
	sli := []string{stra, strb}
	// 1 strings.Join
	fmt.Println(strings.Join(sli, ""))

	// 2 strings.Builder
	var builder strings.Builder
	builder.WriteString("qwertyuiop")
	builder.WriteString("1234567890")
	fmt.Println(builder.String())

	// 3 by.Builder
	var bud bytes.Buffer
	bud.WriteString("qwertyuiop")
	bud.WriteString("1234567890")
	fmt.Println(bud.String())

	// 4 Printf
	printf := fmt.Sprintf("%s%s", "qwertyuiop", "1234567890")
	fmt.Println(printf)

	// 5
	buf := make([]byte, 0)
	buf = append(buf, []byte(stra)...)
	buf = append(buf, []byte(strb)...)
	fmt.Println(string(buf))

	// 6
	fmt.Println(stra + strb)
}
