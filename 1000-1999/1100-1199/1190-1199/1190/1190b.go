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
	a := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
	}
	lost := false
	twice := false
	double := -1
	for k, v := range m {
		if v >= 3 {
			lost = true
			break
		}
		if v >= 2 {
			if k == 0 {
				lost = true
				break
			}
			if twice {
				lost = true
				break
			}
			double = k
			twice = true
		}
	}
	if !lost && twice {
		for i := 0; i < n; i++ {
			if a[i] == double-1 {
				lost = true
				break
			}
		}
		for i := 0; i < n; i++ {
			if a[i] == double {
				a[i]--
				break
			}
		}
	}

	if lost {
		write(f, "cslnb\n")
		return
	}
	sort.Ints(a)

	//	write(f, a, "\n")
	curr := 0
	ans := 0
	for i := 0; i < n; i++ {
		if curr <= a[i] {
			ans += a[i] - curr
			ans %= 2
			curr++
		}
	}
	if twice {
		if ans%2 == 1 {
			write(f, "cslnb\n")
		} else {
			write(f, "sjfnb\n")
		}
	} else {
		if ans%2 == 0 {
			write(f, "cslnb\n")
		} else {
			write(f, "sjfnb\n")
		}
	}

}
