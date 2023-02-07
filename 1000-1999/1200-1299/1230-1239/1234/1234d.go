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
func insert_element(a []int64, x int64) []int64 {
	i := bisect_left(a, x)
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = x
	return a
}
func delete_element(a []int64, x int64) []int64 {
	i := bisect_left(a, x)
	//fmt.Println("searching ", a, x, i)
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]
	return a
}
func has_element(a []int64, l int64, r int64) bool {
	i := bisect_left(a, l)
	if i == len(a) {
		return false
	}
	if a[i] == l {
		return true
	}
	j := bisect_left(a, r)
	if j == len(a) || a[j] == r {
		return true
	}
	if i == j {
		return false
	}
	return true
}

func main() {
	var q int
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s, &q)
	sb := []byte(s)
	sorted := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		sorted[i] = make([]int64, 0)
	}
	for i := 0; i < len(s); i++ {
		sorted[s[i]-'a'] = append(sorted[s[i]-'a'], int64(i))
	}
	//write(f, sorted, "\n")
	for i := 0; i < q; i++ {
		var t, l, r int64
		var st string
		fmt.Fscan(reader, &t)
		if t == 1 {
			fmt.Fscan(reader, &l, &st)
			if sb[l-1] != st[0] {
				re_index := sb[l-1] - 'a'
				//	write(f, "deleting ", l-1, "from", sorted[re_index], "\n")
				sorted[re_index] = delete_element(sorted[re_index], int64(l-1))
				//	write(f, "after delete ", sorted[re_index], "\n")

				sb[l-1] = st[0]
				i_index := st[0] - 'a'
				sorted[i_index] = insert_element(sorted[i_index], int64(l-1))
			}
		} else {
			fmt.Fscan(reader, &l, &r)
			uniq := 0
			for i := 0; i < 26; i++ {
				if has_element(sorted[i], l-1, r-1) {
					//write(f, "searching ", sorted[i], l-1, r-1, "\n")
					uniq++
				}
			}
			write(f, uniq, "\n")
		}
	}

}
