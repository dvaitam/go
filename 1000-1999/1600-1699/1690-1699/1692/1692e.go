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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, s int
		fmt.Fscan(reader, &n, &s)
		a := make([]int, n)
		sm := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			sm += a[i]
		}
		if sm < s {
			write(f, "-1\n")
		} else if sm == s {
			write(f, "0\n")
		} else {
			s = sm - s
			m1 := map[int]int{}
			m2 := map[int]int{}
			sm = 0
			ans := n

			for i := 0; i < n; i++ {
				sm += a[i]
				if a[i] == 1 {
					m1[sm] = i
				}
				if sm == s {
					ans = min(ans, i+1)
				}
			}
			sm = 0
			for i := n - 1; i >= 0; i-- {
				sm += a[i]
				if a[i] == 1 {
					m2[sm] = i
				}
				if sm == s {
					ans = min(ans, n-i)
				}
			}
			//fmt.Println(m1, m2)
			for k, v := range m1 {
				ans = min(ans, v+1+n-m2[s-k])
			}
			write(f, ans, "\n")
		}

	}
}
