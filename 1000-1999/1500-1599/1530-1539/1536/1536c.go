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
func gcd(a int64, b int64) int64 {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
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
		d, k := int64(0), int64(0)
		coff := int64(1000000)
		m := map[int64]int{}
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			if s[i] == 'D' {
				d++
			} else {
				k++
			}
			val := d*coff + k
			if d == 0 {
				val = 1
			} else if k == 0 {
				val = coff
			} else {
				gd := gcd(d, k)
				val = (d/gd)*coff + (k / gd)
			}
			m[val]++
			ans[i] = m[val]
		}
		for i := 0; i < n; i++ {
			write(f, ans[i], " ")
		}
		write(f, "\n")
	}
}
