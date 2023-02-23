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
func split(a []int, b []int, l1 []int, l2 []int, bit int) [][][]int {
	val := 1 << (bit - 1)
	first_ones := make([]int, 0)
	first_zeros := make([]int, 0)
	second_ones := make([]int, 0)
	second_zeros := make([]int, 0)
	for i := 0; i < len(l1); i++ {
		if a[l1[i]]&val == val {

			first_ones = append(first_ones, l1[i])
		} else {
			first_zeros = append(first_zeros, l1[i])
		}
		if b[l2[i]]&val == val {
			second_ones = append(second_ones, l2[i])
		} else {
			second_zeros = append(second_zeros, l2[i])
		}
	}
	ans := make([][][]int, 0)
	if len(first_ones) > 0 {
		ans = append(ans, [][]int{first_ones, second_zeros})
	}
	if len(first_zeros) > 0 {
		ans = append(ans, [][]int{first_zeros, second_ones})

	}
	return ans
}
func possible(a []int, b []int, l1 []int, l2 []int, bit int) bool {
	val := 1 << (bit - 1)
	one := 0
	zero := 0
	for i := 0; i < len(l1); i++ {
		if a[l1[i]]&val == val {
			one++

		}
		if b[l2[i]]&val != val {

			zero++
		}
	}
	return one == zero
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		b := make([]int, n)
		l1 := make([]int, n)
		l2 := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			l1[i] = i
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
			l2[i] = i
		}

		current := [][][]int{{l1, l2}}
		ans := 0
		for bit := 30; bit > 0; bit-- {
			ok := true
			for i := 0; i < len(current); i++ {
				if possible(a, b, current[i][0], current[i][1], bit) {
					continue
				} else {
					ok = false
					break
				}
			}
			if ok {
				// write(f, "ok for bit ", bit, "\n")
				// write(f, current, "\n")
				ans = ans | (1 << (bit - 1))
				next_current := make([][][]int, 0)
				for i := 0; i < len(current); i++ {
					sp := split(a, b, current[i][0], current[i][1], bit)
					for j := 0; j < len(sp); j++ {
						next_current = append(next_current, sp[j])
					}
				}
				current = next_current
				// write(f, "after split", bit, "\n")
				// write(f, current, "\n")

			}

		}
		write(f, ans, "\n")
	}
}
