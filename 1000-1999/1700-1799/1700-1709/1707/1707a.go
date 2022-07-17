// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		low := 0
		high := n
		for low < high {
			mid := (low + high) >> 1
			now := q
			ok := true
			for i := mid; i < n; i++ {
				if now == 0 {
					ok = false
					break
				}
				if a[i] > now {
					now--
				}
			}
			if ok {
				high = mid
			} else {
				low = mid + 1
			}
		}
		for i := 0; i < low; i++ {
			if a[i] <= q {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		for i := low; i < n; i++ {
			fmt.Print(1)
		}
		fmt.Println()
	}
}
