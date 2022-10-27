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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		sort.Slice(a, func(i, j int) bool { return abs(a[i]) <= abs(a[j]) })
		positive_val := -1
		positive_start := false
		positive_first := -1
		negative_start := false
		negative_first := -1
		products := make([]int, 0)

		for i := n - 1; i >= 0; i-- {
			if a[i] >= 0 {
				if positive_val < 0 {
					positive_val = a[i]
				} else {
					if positive_start {
						products = append(products, positive_first*a[i])
						positive_start = false
					} else {
						positive_start = true
						positive_first = a[i]
					}
				}
			} else {
				if negative_start {
					products = append(products, negative_first*a[i])
					negative_start = false
				} else {
					negative_start = true
					negative_first = a[i]
				}
			}
		}
		sort.Ints(products)
		if positive_val >= 0 && len(products) > 1 {
			ll := len(products)
			write(f, int64(positive_val)*int64(products[ll-1])*int64(products[ll-2]), "\n")
		} else {
			write(f, int64(a[0])*int64(a[1])*int64(a[2])*int64(a[3])*int64(a[4]), "\n")
		}
	}
}
