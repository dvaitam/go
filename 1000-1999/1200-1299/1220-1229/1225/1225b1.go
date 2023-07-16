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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k, d int
		fmt.Fscan(reader, &n, &k, &d)
		shows := make([]int, 101)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		subs := 0
		for i := 0; i < d; i++ {
			if shows[a[i]] == 0 {
				subs++
			}
			shows[a[i]]++
		}
		ans := subs
		//write(f, subs, "\n")
		for i := d; i < n; i++ {
			if shows[a[i]] == 0 {
				subs++
			}
			shows[a[i]]++
			shows[a[i-d]]--
			if shows[a[i-d]] == 0 {
				subs--
			}
			//	write(f, subs, "\n")
			ans = min(ans, subs)
		}
		write(f, ans, "\n")
	}
}
