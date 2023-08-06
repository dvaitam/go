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
	var n, m int
	file, _ := os.Open("input.txt")
	//fmt.Println(err)
	ofile, _ := os.Create("output.txt")

	reader := bufio.NewReader(file)
	f := bufio.NewWriter(ofile)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	var A, B, C int
	fmt.Fscan(reader, &A, &B, &C)
	h := make([]int, m)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			h[j] += a[i][j]
			v[i] += a[i][j]
		}
	}
	for i := 1; i < n; i++ {
		v[i] += v[i-1]
	}
	for i := 1; i < m; i++ {
		h[i] += h[i-1]
	}
	mv := map[int]bool{}
	mh := map[int]bool{}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {
			if v[i] == A && v[j] == A+B {
				mv[i*1000+j] = true
			}
			if v[i] == A && v[j] == A+C {
				mv[i*1000+j] = true
			}
			if v[i] == B && v[j] == B+C {
				mv[i*1000+j] = true
			}
			if v[i] == B && v[j] == B+A {
				mv[i*1000+j] = true
			}
			if v[i] == C && v[j] == B+C {
				mv[i*1000+j] = true
			}
			if v[i] == C && v[j] == C+A {
				mv[i*1000+j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := i + 1; j < m-1; j++ {
			if h[i] == A && h[j] == A+B {
				mh[i*1000+j] = true
			}
			if h[i] == A && h[j] == A+C {
				mh[i*1000+j] = true
			}
			if h[i] == B && h[j] == B+C {
				mh[i*1000+j] = true
			}
			if h[i] == B && h[j] == B+A {
				mh[i*1000+j] = true
			}
			if h[i] == C && h[j] == B+C {
				mh[i*1000+j] = true
			}
			if h[i] == C && h[j] == C+A {
				mh[i*1000+j] = true
			}
		}
	}
	if A+B+C == v[n-1] {
		write(f, len(mv)+len(mh), "\n")
	} else {
		write(f, "0\n")
	}
	// if n == 50 && m == 35 {
	// 	write(f, A, B, C, "\n", "\n", v, "\n", mv, "\n")
	// }

}
