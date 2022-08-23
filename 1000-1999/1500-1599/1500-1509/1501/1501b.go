// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		ans := make([]int, n)
		curr := n
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		for i := n - 1; i >= 0; i-- {
			curr = min(curr, i-a[i]+1)
			if curr <= i {
				ans[i] = 1
			} else {
				ans[i] = 0
			}
		}
		for i := 0; i < n; i++ {
			write_int(f, ans[i])
			if i == n-1 {
				write_string(f, "\n")
			} else {
				write_string(f, " ")
			}
		}
	}
}
