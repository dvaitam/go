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
func main() {
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var s string
	fmt.Fscan(reader, &s)
	counts := map[byte]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		counts[s[i]]++
	}
	keys := make([]byte, 0)
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	ans := make([]byte, n)
	half := n / 2
	if n%2 == 1 {
		half++
	}
	i := 0
	for _, k := range keys {
		for j := 0; j < counts[k]/2; j++ {
			ans[i] = k
			i++
		}
		counts[k] -= (counts[k] / 2) * 2
	}
	for i < half {
		for _, k := range keys {
			if counts[k] == 0 {
				continue
			}
			ans[i] = k
			i++
		}
	}
	if n%2 == 0 {
		sort.Slice(ans[:half], func(i, j int) bool { return ans[i] < ans[j] })

	} else {
		sort.Slice(ans[:half-1], func(i, j int) bool { return ans[i] < ans[j] })

	}
	for i := 0; i < half; i++ {
		ans[n-i-1] = ans[i]
	}
	write(f, string(ans), "\n")

}
