/*
進行實驗以比較無效率版與使用 strings.Join 的版本
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// +=
	loopTimes := 10000
	var result string
	start := time.Now()
	for i := 1; i < loopTimes; i++ {
		result += strconv.FormatInt(int64(i), 10) + " "
	}
	elapsed := time.Since(start)
	fmt.Println("+= took ", elapsed)

	// strings.join
	result = ""
	start = time.Now()
	s := []string{}
	for i := 1; i < loopTimes; i++ {
		s = append(s, strconv.FormatInt(int64(i), 10))
	}
	strings.Join(s, " ")
	elapsed = time.Since(start)
	fmt.Println("strings.Join took ", elapsed)
}
