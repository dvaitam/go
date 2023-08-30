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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	write(f, 8+7*n, "\n")
	lower := make([][]int, 8)
	lower[0] = []int{0, 0}
	lower[1] = []int{0, 1}
	lower[2] = []int{0, 2}
	lower[3] = []int{1, 2}
	lower[4] = []int{2, 2}
	lower[5] = []int{2, 1}
	lower[6] = []int{2, 0}
	lower[7] = []int{1, 0}
	upper := make([][]int, 7)
	upper[0] = []int{2, 3}
	upper[1] = []int{2, 4}
	upper[2] = []int{3, 4}
	upper[3] = []int{4, 4}
	upper[4] = []int{4, 3}
	upper[5] = []int{4, 2}
	upper[6] = []int{3, 2}
	for i := 0; i < 8; i++ {
		write(f, lower[i][0], lower[i][1], "\n")
	}
	up := true
	for i := 0; i < n; i++ {
		if up {
			for j := 0; j < 7; j++ {
				write(f, upper[j][0], upper[j][1], "\n")
				upper[j][0] += 4
			}
		} else {
			for j := 0; j < 8; j++ {
				if j == 2 {
					continue
				}
				lower[j][0] += 4
				write(f, lower[j][0], lower[j][1], "\n")
			}
		}
		up = !up
	}

}
