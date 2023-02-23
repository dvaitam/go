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

var swaps int

func quicksort(a []int, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := partition(a, lo, hi)
	quicksort(a, lo, p-1)
	quicksort(a, p+1, hi)
}
func partition(a []int, lo int, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
			swaps++
		}
	}
	i++
	a[i], a[hi] = a[hi], a[i]
	swaps++
	return i
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	//first_answer := 0
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		p := make([]int, n)
		ans := 0
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
			m[p[i]] = i + 1
		}
		quicksort(p, 0, n-1)
		write(f, swaps, p, "\n")
		// if t == 219 && first_answer == 114 {
		// 	fmt.Println(p)
		// }
		adjacent := false
		for i := 0; i < n; i++ {
			if p[i] == i+1 {
				continue
			} else {
				if !adjacent && i+1 < n {
					if p[i] == i+2 && p[i+1] == i+1 {
						adjacent = true
						i++
					} else if p[i] == i+2 {
						if p[i+1] != i+1 {
							si, sj := i+1, m[i+1]-1
							p[si], p[sj] = p[sj], p[si]
							m[p[si]] = si + 1
							m[p[sj]] = sj + 1
							ans++

						}
						adjacent = true
						i += 1
					} else if p[i+1] == i+1 {
						if p[i] != i+2 {
							si, sj := i, m[i+2]-1
							p[si], p[sj] = p[sj], p[si]
							m[p[si]] = si + 1
							m[p[sj]] = sj + 1
							ans++
						}
						adjacent = true
						i += 1

					} else if m[i+2] == p[i] && m[i+1] == p[i+1] {
						if p[i] != i+2 {
							si, sj := i, m[i+2]-1
							p[si], p[sj] = p[sj], p[si]
							m[p[si]] = si + 1
							m[p[sj]] = sj + 1
							ans++
						}
						if p[i+1] != i+1 {
							si, sj := i+1, m[i+1]-1
							p[si], p[sj] = p[sj], p[si]
							m[p[si]] = si + 1
							m[p[sj]] = sj + 1
							ans++

						}
						adjacent = true
						i += 1
					} else {
						si, sj := i, m[i+1]-1

						p[si], p[sj] = p[sj], p[si]
						m[p[si]] = si + 1
						m[p[sj]] = sj + 1
						ans++
					}

				} else {
					si, sj := i, m[i+1]-1

					p[si], p[sj] = p[sj], p[si]
					m[p[si]] = si + 1
					m[p[sj]] = sj + 1
					ans++
				}

			}
			write(f, p, "\n")
			//write(f, m, "\n")
		}
		if !adjacent {
			ans++
		}
		// if t == 1 {
		// 	first_answer = ans
		// }

		write(f, ans, "\n")

	}
}
