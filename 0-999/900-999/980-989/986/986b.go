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
	m := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]] = i
	}
	start := 1
	ans := 0
	for start <= n {
		if m[start] != start-1 {
			i, j := start-1, m[start]
			a[i], a[j] = a[j], a[i]
			m[a[i]], m[a[j]] = i, j
			ans++
		}
		start++
	}
	//write(f, ans, a, "\n")
	if (3*n-ans)%2 == 0 {
		write(f, "Petr\n")
	} else {
		write(f, "Um_nik\n")
	}

}
