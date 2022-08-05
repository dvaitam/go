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
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var m int
		fmt.Fscan(reader, &m)
		a := make([][]int, 2)
		a[0] = make([]int, m)
		a[1] = make([]int, m)
		for i := 0; i < 2; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(reader, &a[i][j])
			}
		}
		ans1 := 0
		for j := 1; j < m; j++ {
			ans1 = max(ans1+1, a[0][j]+1)
		}
		for j := m - 1; j >= 0; j-- {
			ans1 = max(ans1+1, a[1][j]+1)
		}

		ans2 := max(1, a[1][0])
		for j := 1; j < m; j++ {
			ans2 = max(ans2+1, a[1][j]+1)
		}
		for j := m - 1; j > 0; j-- {
			ans2 = max(ans2+1, a[0][j]+1)
		}
		if m == 2 {
			fmt.Println(min(ans1, ans2))
		} else {
			ans3 := max(1, a[1][0])
			ans3 = max(ans3+1, a[1][1]+1)
			ans3 = max(ans3+1, a[0][1]+1)
			for j := 2; j < m; j++ {
				ans3 = max(ans3+1, a[0][j]+1)
			}
			for j := m - 1; j > 1; j-- {
				ans3 = max(ans3+1, a[1][j]+1)
			}
			fmt.Println(min(min(ans1, ans2), ans3))
		}

	}
}
