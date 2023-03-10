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
func get_agg(p []float64) float64 {
	n := len(p)
	pr := float64(1)
	pa := float64(0)
	for i := 0; i < n; i++ {
		pr *= (1 - p[i])
		pa += 1 / (1 - p[i])
	}
	return pr * (pa - float64(n))
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	p := make([]float64, n)
	has_one := false
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
		if p[i] == 1 {
			has_one = true
		}

	}
	if has_one {
		write(f, "1.0\n")
	} else {
		ans := float64(0)
		sort.Float64s(p)
		for i := 0; i < n; i++ {
			ans = max(ans, get_agg(p[i:]))
		}
		write(f, ans, "\n")
	}

}
