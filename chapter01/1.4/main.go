/*
修改 dup2 以輸出重複行的檔案名稱
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file name, line text, count
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "from command line", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for fileName, line := range counts {
		for l, n := range line {
			if n > 1 {
				fmt.Printf("file: %s\tline: %s\tcount: %d\n", fileName, l, n)
			}
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[fileName] == nil {
			counts[fileName] = make(map[string]int)
		}
		counts[fileName][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
