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
type Object struct {
	v, i int
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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		a := make([]int, n)
		d := make([]int, n)

		for i := 0; i < n; i++ {
			a[i], _ = strconv.Atoi(s[i : i+1])
		}

		m := map[int]int{}
		sm := 0
		for i := 0; i < n; i++ {

			sm += a[i] - 1
			if m[sm] > 0 {
				d[i] = d[m[sm]-1] + 1
			} else if sm == 0 {
				d[i] = 1
			}
			m[sm] = i + 1

		}
		//write(f, m, "\n")
		//write(f, d, "\n")

		ans := int64(0)
		for i := 0; i < n; i++ {
			ans += int64(d[i])
		}
		write(f, ans, "\n")
	}
}
