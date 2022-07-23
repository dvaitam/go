// 1471A. Strange Partition
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, x int
		fmt.Fscan(reader, &n, &x)
		a := make([]int, n)
		sum := int64(0)
		max := int64(0)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			sum += int64(a[i])
			if a[i]%x == 0 {
				max += int64(a[i] / x)
			} else {
				max += int64((a[i] / x) + 1)
			}
		}
		min := sum / int64(x)
		if sum%int64(x) != 0 {
			min++
		}
		fmt.Println(min, max)
	}
}
