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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	mi := 1000000000
	mx := 0
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		mi = min(mi, a[i])
		mx = max(mx, a[i])
		m[a[i]]++
	}
	if len(m) > 3 {
		write(f, "-1\n")
	} else {
		other := -1
		for k := range m {
			if k != mi && k != mx {
				other = k
			}
		}
		if other != -1 {
			if 2*other == mi+mx {
				write(f, mx-other, "\n")
			} else {
				write(f, "-1\n")
			}
		} else {
			if (mi+mx)%2 == 0 {
				write(f, min(((mi+mx)/2)-mi, mx-mi), "\n")
			} else {
				write(f, mx-mi, "\n")
			}
		}
	}

}
