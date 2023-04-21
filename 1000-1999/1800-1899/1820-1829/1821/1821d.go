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

type Segment struct {
	l, r int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	first_n := -1
	for t := 1; t <= T; t++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		if first_n == -1 {
			first_n = n

		}
		l := make([]int, n)
		r := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &l[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &r[i])
		}
		// if first_n == 20 {
		// 	if t == 191 {
		// 		fmt.Println(n, k, l, r)
		// 	} else {
		// 		continue
		// 	}
		// }
		ones := make([]int, 0)
		ones_not_used := make([]int, 0)
		others := make([]Segment, 0)
		for i := 0; i < n; i++ {
			if l[i] == r[i] {
				ones_not_used = append(ones_not_used, l[i])
			} else if k > 0 {
				if r[i]-l[i]+1 <= k {
					others = append(others, Segment{l: l[i], r: r[i]})
					k -= r[i] - l[i] + 1
				} else {
					others = append(others, Segment{l: l[i], r: r[i] - ((r[i] - l[i] + 1) - k)})
					k = 0
				}
			} else {
				break
			}
		}
		for k > 0 && len(ones_not_used) >= k {
			ones = ones_not_used[:k]
			ones_not_used = ones_not_used[k:]
			k = 0
		}
		if k > 0 {
			write(f, "-1\n")
			continue
		}
		if len(others) == 0 {
			ending := 0

			if len(ones) > 0 {
				ending = max(ending, ones[len(ones)-1])
			}
			write(f, ending+2*len(ones), "\n")
			continue
		}
		ending := others[len(others)-1].r
		if len(ones) > 0 {
			ending = max(ending, ones[len(ones)-1])
		}
		ans := ending + 2*(len(ones)+len(others))
		curr := len(others) - 1
		start := 0
		for curr >= 0 && start < len(ones_not_used) {
			if others[curr].l < ones_not_used[start] {
				break
			}
			start += others[curr].r - others[curr].l + 1
			if start-1 >= len(ones_not_used) {
				break
			}
			new_ending := ones_not_used[start-1]
			if curr-1 >= 0 {
				new_ending = max(new_ending, others[curr-1].r)
			}
			ans = min(ans, new_ending+curr*2+len(ones)*2+start*2)
			curr--
		}
		write(f, ans, "\n")
	}
}
