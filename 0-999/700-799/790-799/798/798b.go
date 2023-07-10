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
func index(s1 string, s2 string) int {
	for i := 0; i < len(s1); i++ {
		ok := true
		for j := 0; j < len(s1); j++ {
			if s1[j] != s2[(i+j)%len(s1)] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	d := make([]int, n)
	for i := 1; i < n; i++ {
		d[i] = index(s[0], s[i])
	}
	//write(f, d, "\n")
	ok := true
	for i := 0; i < n; i++ {
		if d[i] == -1 {
			ok = false
			break
		}
	}
	if ok {
		ans := 0
		for i := 1; i < n; i++ {
			ans += d[i]
		}
		for i := 1; i < n; i++ {
			tmp := 0
			for j := 0; j < n; j++ {
				// if d[j] > d[i] {
				// 	tmp += d[j] - d[i]
				// } else if d[j] < d[i] {
				// 	tmp += len(s[0]) - (d[i] - d[j])
				// }
				tmp += index(s[i], s[j])
			}
			ans = min(ans, tmp)
		}
		write(f, ans, "\n")
	} else {
		write(f, "-1\n")
	}

}
