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
	a := make([]int64, 2000001)
	a[1] = 1
	start, end := 1, 1
	curr := 1
	for true {
		start, end = end+1, end+1+curr
		prev_start := ((curr - 1) * curr) / 2
		prev_start++
		prev_end := prev_start + curr - 1
		if start > 1000000 {
			break
		}
		//write(f, start, end, prev_start, prev_end, "\n")
		//write(f, prev_start, prev_end, "\n")
		for i := start; i <= end; i++ {
			a[i] = int64(i * i)
			if i == start {
				a[i] += a[prev_start]
			} else if i == end {
				a[i] += a[prev_end]
			} else {
				sub := i - start + ((curr-2)*(curr-1))/2
				a[i] += a[prev_start+i-start-1]
				a[i] += a[prev_start+i-start]
				// if i == 9 {
				// 	write(f, "a is ", a[i], prev_start+i-start-1, prev_start+i-start, "\n")
				// 	write(f, "sub is", sub, curr, "\n")
				// }
				if sub > 0 {
					a[i] -= a[sub]
				}
			}
		}
		curr++
	}
	//	write(f, a[:20], "\n")
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		write(f, a[n], "\n")
	}
}
