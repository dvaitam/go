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
	var n, I int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &I)
	a := make([]int, n)
	counter := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		counter[a[i]]++
	}
	uniq := make([]int, 0)
	for k := range counter {
		uniq = append(uniq, k)
	}
	sort.Ints(uniq)
	curr := len(uniq)
	total := I * 8
	mbits := total / n
	acc := 1 << mbits
	if acc >= curr || mbits > 30 {
		write(f, "0\n")
	} else {
		sac := 0
		for i := acc; i < len(uniq); i++ {
			sac += counter[uniq[i]]
		}
		ans := sac
		left, right := 0, acc
		for right < len(uniq) {
			sac = sac - counter[uniq[right]] + counter[uniq[left]]
			ans = min(ans, sac)
			left++
			right++
		}
		write(f, ans, "\n")
	}

}
