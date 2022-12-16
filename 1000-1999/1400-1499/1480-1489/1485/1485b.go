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
	var n, q, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &q, &k)
	a := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		if i+1 < n {
			d[i] += a[i+1] - a[i] - 1
		}
	}
	//fmt.Println(d)
	for i := 1; i < n; i++ {
		d[i] += d[i-1]
	}
	//fmt.Println(d)
	var l, r int
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &l, &r)
		if l == r {
			write(f, k-1, "\n")
		} else {
			l--
			r--
			if a[r] <= k {
				if l == 0 {
					write(f, 2*d[r-1]+(a[0]-1)+(k-a[r]), "\n")
				} else {
					write(f, 2*(d[r-1]-d[l-1])+(a[l]-1)+(k-a[r]), "\n")
				}
			} else {
				write(f, "0\n")
			}
		}

	}

}
