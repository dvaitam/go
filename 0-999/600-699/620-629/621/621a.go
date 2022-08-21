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
	var curr int64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	sum := int64(0)
	odds_count := 0
	min_odd := int64(1000000000)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &curr)
		if curr%2 == 1 {
			min_odd = min(min_odd, curr)
			odds_count++
		}
		sum += curr
	}
	if odds_count%2 == 1 {
		sum -= min_odd
	}
	fmt.Println(sum)
}
