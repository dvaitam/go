// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	min_index := -1
	min_val := 1000000001
	duplicate := false
	curr := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &curr)
		if curr < min_val {
			min_val = curr
			min_index = i
			duplicate = false
		} else if curr <= min_val {
			duplicate = true
		}
	}
	if duplicate {
		fmt.Println("Still Rozdil")
	} else {
		fmt.Println(min_index)
	}

}
