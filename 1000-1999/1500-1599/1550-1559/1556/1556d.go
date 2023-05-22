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
func decode(anb bool, anc bool, bnc bool, aob bool, aoc bool, boc bool) (bool, bool, bool) {
	if anb {
		if anc {
			return true, true, true
		} else {
			return true, true, false
		}
	} else {
		if aob {
			if aoc {
				if anc {
					return true, false, true
				} else {
					if boc {
						return false, true, true
					} else {
						return true, false, false
					}
				}

			} else {
				return false, true, false
			}
		} else {
			if aoc {
				return false, false, true
			} else {
				return false, false, false
			}
		}
	}
}
func main() {
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	var anb, anc, bnc, aob, aoc, boc int
	fmt.Println("and 1 2")
	fmt.Fscan(reader, &anb)
	fmt.Println("and 1 3")
	fmt.Fscan(reader, &anc)
	fmt.Println("and 2 3")
	fmt.Fscan(reader, &bnc)
	fmt.Println("or 1 2")
	fmt.Fscan(reader, &aob)
	fmt.Println("or 1 3")
	fmt.Fscan(reader, &aoc)
	fmt.Println("or 2 3")
	fmt.Fscan(reader, &boc)
	for j := 0; j < 30; j++ {
		r := 1 << j
		ai, bi, ci := decode(anb&r == r, anc&r == r, bnc&r == r, aob&r == r, aoc&r == r, boc&r == r)
		if ai {
			a[0] += r
		}
		if bi {
			a[1] += r
		}
		if ci {
			a[2] += r
		}
	}
	for i := 3; i < n; i++ {
		var ni, oi int
		fmt.Println("and 1 ", i+1)
		fmt.Fscan(reader, &ni)
		fmt.Println("or 1 ", i+1)
		fmt.Fscan(reader, &oi)
		for j := 0; j < 30; j++ {
			r := 1 << j
			if a[0]&r == r {
				a[i] += ni & r
			} else {
				a[i] += oi & r
			}
		}
	}
	//fmt.Println(a)
	sort.Ints(a)
	fmt.Println("finish ", a[k-1])
}
