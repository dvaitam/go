// 567A. Lineland Mail
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Println(a[i+1]-a[i], a[n-1]-a[i])
		} else if i == n-1 {
			fmt.Println(a[i]-a[i-1], a[i]-a[0])
		} else {
			fmt.Println(min(a[i]-a[i-1], a[i+1]-a[i]), max(a[i]-a[0], a[n-1]-a[i]))
		}
	}
}
