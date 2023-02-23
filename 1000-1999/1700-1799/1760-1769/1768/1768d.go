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
func inverts(a []int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				ans++
			}
		}
	}
	return ans
}

var steps int
var zero_steps int

func backtrack(a []int, count int) {
	inc := inverts(a)
	if inc <= 1 {
		if inc == 0 {
			zero_steps = min(zero_steps, count)
		}
		if inc == 1 {

			steps = min(steps, count)
		}
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != i+1 {
			si, sj := i, a[i]-1
			a[si], a[sj] = a[sj], a[si]
			backtrack(a, count+1)
			a[si], a[sj] = a[sj], a[si]
		}
	}

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
		m := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
			m[p[i]] = i + 1
		}
		steps = n
		zero_steps = n
		backtrack(p, 0)
		write(f, min(steps, zero_steps+1), "\n")

	}
}
