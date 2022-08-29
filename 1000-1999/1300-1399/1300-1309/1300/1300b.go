// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
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
		a := make([]int, 2*n)
		for i := 0; i < 2*n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		sort.Ints(a)
		ans := 1000000000
		for i := 0; i < 2*n; i++ {
			if i < n {
				ans = min(ans, abs(a[i]-a[n]))
			} else {
				ans = min(ans, abs(a[i]-a[n-1]))
			}
		}
		fmt.Println(ans)
	}
}
