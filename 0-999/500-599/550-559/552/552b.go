// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func pow(a int, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := strconv.Itoa(n)
	ans := 0
	for d := 1; d < len(s); d++ {
		ans += 9 * pow(10, d-1) * d
	}
	start := 1 * pow(10, len(s)-1)
	ans += (n - start + 1) * len(s)
	write(f, ans, "\n")
}
