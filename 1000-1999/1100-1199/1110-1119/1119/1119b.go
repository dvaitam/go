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
func min_height(a []int) int64 {
	b := make([]int, 0)
	for i := 0; i < len(a); i++ {
		b = append(b, a[i])
	}
	sort.Ints(b)
	if len(b)%2 == 0 {
		ans := int64(0)
		for i := 1; i < len(b); i += 2 {
			ans += int64(b[i])
		}
		return ans
	} else {
		ans1 := int64(b[len(b)-1])
		for i := 1; i < len(b); i += 2 {
			ans1 += int64(b[i])
		}
		ans2 := int64(0)
		for i := 0; i < len(b); i += 2 {
			ans2 += int64(b[i])
		}
		//fmt.Println(ans1, ans2, "\n")
		return min(ans1, ans2)
	}
}
func main() {
	var n int
	var h int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &h)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	ans := 0
	for k := n; k > 0; k-- {
		min_h := min_height(a[:k])
		//write(f, min_h, "\n")
		if min_h <= h {
			ans = k
			break
		}
	}
	write(f, ans, "\n")

}
