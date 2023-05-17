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

type Segment struct {
	l, r, q int
}

func update(a []int, i int, l int, r int, l1 int, r1 int, q int) {
	if l1 == l && r == r1 {
		a[i] |= q
		//fmt.Println("updating ", i, a[i])
	} else {
		if r1 <= (l+r)/2 {
			update(a, 2*i+1, l, (l+r)/2, l1, r1, a[i]|q)
		} else if l1 >= 1+(l+r)/2 {
			update(a, 2*i+2, 1+(l+r)/2, r, l1, r1, a[i]|q)
		} else {
			update(a, 2*i+1, l, (l+r)/2, l1, (l+r)/2, a[i]|q)
			update(a, 2*i+2, 1+(l+r)/2, r, 1+(l+r)/2, r1, a[i]|q)
		}
	}
}
func verify(a []int, i int, l int, r int, l1 int, r1 int, q int) bool {
	if l1 == l && r == r1 {
		return a[i] == q
	} else {
		if r1 <= (l+r)/2 {
			return verify(a, 2*i+1, l, (l+r)/2, l1, r1, q)
		} else if l1 >= 1+(l+r)/2 {
			return verify(a, 2*i+2, 1+(l+r)/2, r, l1, r1, q)
		} else {
			return verify(a, 2*i+1, l, (l+r)/2, l1, (l+r)/2, q) || verify(a, 2*i+2, 1+(l+r)/2, r, 1+(l+r)/2, r1, q)
		}
	}
}
func val(a []int, i int, l int, r int, j int) int {
	if j == l && j == r {
		return a[i]
	}
	if j <= (l+r)/2 {
		return a[i] | val(a, 2*i+1, l, (l+r)/2, j)
	} else {
		return a[i] | val(a, 2*i+2, 1+(l+r)/2, r, j)
	}
}

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	from := make([]int, m)
	to := make([]int, m)
	num := make([]int, m)
	s := make([][]int, n+1)
	a := make([][]bool, n+1)
	sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		s[i] = make([]int, 30)
		a[i] = make([]bool, 30)
		sum[i] = make([]int, 30)
	}
	//segments := make([]Segment, m)
	for i := 0; i < m; i++ {
		var l, r, q int
		fmt.Fscan(reader, &l, &r, &q)
		l--
		r--
		from[i] = l
		to[i] = r
		num[i] = q
		for j := 0; j < 30; j++ {
			if q&(1<<j) != 0 {
				s[l][j]++
				s[r+1][j]--
			}
		}
		//segments[i] = Segment{l - 1, r - 1, q}
	}
	// a := make([]int, 262144)
	// for i := 0; i < m; i++ {
	// 	update(a, 0, 0, 131071, segments[i].l, segments[i].r, segments[i].q)
	// }
	// ok := true
	// for i := 0; i < m; i++ {
	// 	ok = verify(a, 0, 0, 131071, segments[i].l, segments[i].r, segments[i].q)
	// 	if !ok {
	// 		break
	// 	}
	// }
	// if ok {
	// 	write(f, "YES\n")
	// 	for i := 0; i < n; i++ {
	// 		write(f, val(a, 0, 0, 131071, i), " ")
	// 	}
	// 	write(f, "\n")
	// } else {
	// 	write(f, "NO\n")
	// }

	for j := 0; j < 30; j++ {
		bal := 0
		sum[0][j] = 0
		for i := 0; i < n; i++ {
			bal += s[i][j]
			a[i][j] = bal > 0
			if a[i][j] {
				sum[i+1][j] = sum[i][j] + 1
			} else {
				sum[i+1][j] = sum[i][j]
			}
		}
	}
	for k := 0; k < m; k++ {
		for j := 0; j < 30; j++ {
			if (num[k] & (1 << j)) == 0 {
				get := sum[to[k]+1][j] - sum[from[k]][j]
				need := (to[k] + 1) - (from[k])
				if get == need {
					write(f, "NO\n")
					return
				}
			}
		}
	}
	write(f, "YES\n")
	for i := 0; i < n; i++ {
		res := 0
		for j := 0; j < 30; j++ {
			if a[i][j] {
				res += (1 << j)
			}
		}
		if i > 0 {
			write(f, " ")
		}
		write(f, res)
	}
	write(f, "\n")
}
