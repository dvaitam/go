// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		ans := 0

		rows := n / 2
		cols := n / 2
		if n%2 == 1 {
			cols++
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				count := 0
				if a[i][j] == '1' {
					count++
				}
				if a[j][n-1-i] == '1' {
					count++
				}
				if a[n-1-i][n-1-j] == '1' {
					count++
				}
				if a[n-1-j][i] == '1' {
					count++
				}
				ans += min(count, 4-count)
				//				fmt.Println(count, 4-count)
			}
		}

		fmt.Println(ans)
	}
}
