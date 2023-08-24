// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, n)
	sm := int64(0)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		sm += a[i]
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	mx := int64(31624)
	if n == 4 {
		mx = 1000
	} else if n == 5 {
		mx = 200
	} else if n == 6 {
		mx = 100
	} else if n == 7 {
		mx = 32
	} else if n == 8 || n == 9 {
		mx = 20
	} else if n >= 10 {
		mx = 10
	}
	ans := sm
	for i := int64(1); i <= mx; i++ {
		curr := int64(1)
		sm := int64(0)
		for j := 0; j < n; j++ {
			if curr < 0 {
				break
			}
			sm += abs(curr - a[j])
			curr = curr * i
		}
		if curr < 0 {
			break
		}
		//	write(f, sm, "\n")
		if sm >= 0 {
			ans = min(ans, sm)
		}
	}
	write(f, ans, "\n")

}
