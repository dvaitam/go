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
	var k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &k)
	var s string
	fmt.Fscan(reader, &s)
	n := len(s)
	a := make([][]int, 0)
	start := 1 << k
	for start > 0 {
		a = append(a, make([]int, start))
		start = start >> 1
	}
	for i := 0; i < 1<<k; i++ {
		a[0][i] = 1
	}
	j := 1
	prev := 0
	for i := 0; i < n; i++ {
		if i-prev == len(a[j]) {
			prev += len(a[j])
			j++
		}
		if s[i] == '?' {
			a[j][i-prev] = a[j-1][2*(i-prev)] + a[j-1][2*(i-prev)+1]
		} else if s[i] == '0' {
			a[j][i-prev] = a[j-1][2*(i-prev)]
		} else {
			a[j][i-prev] = a[j-1][2*(i-prev)+1]
		}
	}
	//write(f, a, "\n")
	var q int
	fmt.Fscan(reader, &q)
	sb := []byte(s)
	for t := 1; t <= q; t++ {
		var index int
		var ch string
		fmt.Fscan(reader, &index, &ch)
		index--
		sb[index] = ch[0]
		j := 1
		prefix := 0
		for j < len(a) {
			if index >= len(a[j]) {
				index -= len(a[j])
				prefix += len(a[j])
				j++
			} else {
				if sb[prefix+index] == '?' {
					a[j][index] = a[j-1][2*index] + a[j-1][2*index+1]
				} else if sb[prefix+index] == '0' {
					a[j][index] = a[j-1][2*index]
				} else {
					a[j][index] = a[j-1][2*index+1]
				}
				prefix += len(a[j])
				index = index / 2
				j++
			}
		}
		write(f, a[len(a)-1][0], "\n")
	}
}
