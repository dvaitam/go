// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	b := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	a := make([]int64, n)
	a[0] = b[0]
	mx := a[0]
	for i := 1; i < n; i++ {
		a[i] = b[i] + mx
		mx = max(a[i], mx)
	}
	fmt.Print(a[0])
	for i := 1; i < n; i++ {
		fmt.Print(" ", a[i])
	}
	fmt.Println()

}
