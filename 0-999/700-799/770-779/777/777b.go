// 00
package main

import (
	"bufio"
	"fmt"
	"os"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	var n int
	var s1, s2 string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &s1, &s2)
	a := make([]int, 10)
	b := make([]int, 10)
	for i := 0; i < n; i++ {
		a[int(s1[i]-'0')]++
		b[int(s2[i]-'0')]++
	}
	mi := make([]int, 10)
	for i := 0; i <= 9; i++ {
		curr := b[i]
		if curr == 0 {
			continue
		}
		for j := i; j >= 0; j-- {
			if curr < a[j]-mi[j] {
				mi[j] += curr
				curr = 0
				break
			} else {
				curr -= a[j] - mi[j]
				mi[j] = a[j]
			}
		}
	}
	mx := make([]int, 10)
	for i := 0; i <= 9; i++ {
		curr := b[i]
		if curr == 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if curr < a[j]-mx[j] {
				mx[j] += curr
				curr = 0
				break
			} else {
				curr -= a[j] - mx[j]
				mx[j] = a[j]
			}
		}
	}
	mi_ans, mx_ans := n, 0
	for i := 0; i <= 9; i++ {
		mi_ans -= mi[i]
	}
	for i := 0; i <= 9; i++ {
		mx_ans += mx[i]
	}
	write(f, mi_ans, "\n", mx_ans, "\n")

}
