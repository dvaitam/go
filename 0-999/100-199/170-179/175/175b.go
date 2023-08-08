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
	m := map[string]int{}
	for i := 0; i < n; i++ {
		var s string
		var c int
		fmt.Fscan(reader, &s, &c)
		m[s] = max(m[s], c)
	}
	all := make([]string, 0)
	for s := range m {
		all = append(all, s)
	}
	sort.Slice(all, func(i, j int) bool { return m[all[i]] < m[all[j]] })
	write(f, len(all), "\n")
	for i := 0; i < len(all); i++ {
		rem := 0
		for j := i + 1; j < len(all); j++ {
			if m[all[j]] > m[all[i]] {
				rem++
			}
		}
		if 2*rem > len(all) {
			write(f, all[i], " noob\n")
		} else if 5*rem > len(all) {
			write(f, all[i], " random\n")
		} else if 10*rem > len(all) {
			write(f, all[i], " average\n")
		} else if 100*rem > len(all) {
			write(f, all[i], " hardcore\n")
		} else {
			write(f, all[i], " pro\n")
		}
	}
}
