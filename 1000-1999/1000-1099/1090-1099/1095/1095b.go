// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	if n <= 2 {
		fmt.Println(0)
	} else {
		fmt.Println(min(a[n-1]-a[1], a[n-2]-a[0]))
	}
}
