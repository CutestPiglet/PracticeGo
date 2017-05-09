/*
修改 echo 程式以分行輸出每個參數的索引與值
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Printf("index = %d, value = %s\n", index, value)
	}
}
