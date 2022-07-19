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
func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		h := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &h[i])
		}
		ans := int64(0)
		for i := 1; i < n-1; i += 2 {
			mx := max(h[i-1], h[i+1])
			if mx >= h[i] {
				ans += int64((mx - h[i] + 1))
			}
		}
		if n%2 == 1 {
			fmt.Println(ans)
		} else {
			cool := make([]int, n)
			for i := 1; i < n-1; i++ {
				mx := max(h[i-1], h[i+1])
				if mx >= h[i] {
					cool[i] = (mx - h[i] + 1)
				}
			}
			// fmt.Println(cool)
			ans1 := int64(0)
			for i := 1; i < n-1; i += 2 {
				ans1 += int64(cool[i])
			}
			ans := ans1
			//fmt.Println("starting is ", ans1)
			for i := n - 3; i > 0; i -= 2 {
				ans1 += int64(cool[i+1] - cool[i])
				ans = min(ans, ans1)
				// fmt.Println("is ", i, ans1)
			}
			fmt.Println(ans)
		}

	}
}
