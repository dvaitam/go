// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	l, r int64
}

func bisect_right(a []int64, x int64) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if x < a[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
func bisect_left(a []int64, x int64) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, c, q int
		var s string
		fmt.Fscan(reader, &n, &c, &q, &s)
		pairs := make([]Pair, c)
		csum := make([]int64, 0)
		csum = append(csum, int64(len(s)))
		for i := 0; i < c; i++ {
			var l, r int64
			fmt.Fscan(reader, &l, &r)
			pairs[i] = Pair{l: l - 1, r: r - 1}
			csum = append(csum, csum[len(csum)-1]+(r-l+1))
		}

		// fmt.Println(csum)
		var k int64
		for i := 0; i < q; i++ {
			fmt.Fscan(reader, &k)
			k--
			index := bisect_right(csum, k)
			for index != 0 {
				// fmt.Println("searching for ", k, index)
				k = pairs[index-1].l + (k - csum[index-1])
				index = bisect_right(csum, k)
			}
			fmt.Printf("%c\n", s[k])
		}

	}
}
