/*
撰寫類似 cf 的通用單位轉換程式，從命令列參數或沒有參數時從標準輸入讀取數字，將數字轉換為攝氏與華氏、英呎與公尺、英鎊與公斤等單元
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"chapter02/2.2/lengthconv"
	"chapter02/2.2/tempconv"
	"chapter02/2.2/weightconv"
)

func convert(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	// temperature
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
	// length
	feet := lengthconv.Feet(t)
	m := lengthconv.Meters(t)
	fmt.Printf("%s = %s, %s = %s\n",
		feet, lengthconv.FToM(feet), m, lengthconv.MToF(m))
	// weight
	p := weightconv.Pound(t)
	k := weightconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n",
		p, weightconv.PToK(p), k, weightconv.KToP(k))
}

func main() {
	// command line arguments
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
		return
	}

	// stdin
	for true {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			break
		}
		convert(arg)
	}
}
