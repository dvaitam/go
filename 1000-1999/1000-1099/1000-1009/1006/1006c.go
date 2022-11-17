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
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	var n int
	fmt.Fscan(reader, &n)
	d := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}
	ans := int64(0)
	left_sum := int64(0)
	right_sum := int64(0)
	i := 0
	j := n - 1
	for i <= j {
		if left_sum == right_sum {
			ans = max(ans, left_sum)
			left_sum += d[i]
			i++
		} else if left_sum < right_sum {
			left_sum += d[i]
			i++
		} else if right_sum < left_sum {
			right_sum += d[j]
			j--
		}
		//write(f, left_sum, right_sum, i, j, "\n")
	}
	if left_sum == right_sum {
		ans = max(ans, left_sum)
	}
	write(f, ans, "\n")

}
