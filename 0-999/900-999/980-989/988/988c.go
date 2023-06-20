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
	var k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &k)
	mk := map[int]int{}
	mj := map[int]int{}
	found := false
	for i := 0; i < k; i++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		sm := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[j])
			sm += a[j]
		}
		for j := 0; j < n; j++ {
			curr := sm - a[j]
			if mk[curr] != 0 && mk[curr] != i+1 {
				write(f, "YES\n")
				write(f, mk[curr], mj[curr], "\n")
				write(f, i+1, j+1, "\n")
				found = true
				break
			}
			mk[curr] = i+1
			mj[curr] = j+1
		}
		if found {
			break
		}
	}
	if !found {
		write(f, "NO\n")
	}
}
