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
		mx := make([]int, n)
		b := make([]byte, n)
		if a[n-1] <= q {
			mx[n-1] = a[n-1]
		}
		for i := n - 2; i >= 0; i-- {
			if a[i] <= q {
				mx[i] = max(a[i], mx[i+1])
			} else {
				mx[i] = mx[i+1]
			}
		}
		//fmt.Println(mx)
		for i := 0; i < n; i++ {
			if q >= a[i] {
				b[i] = '1'
			} else {
				if q > mx[i] {
					b[i] = '1'
					q--
				} else {
					b[i] = '0'
				}
			}
		}

		fmt.Println(string(b))
	}
}
