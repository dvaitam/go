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

type Ending struct {
	pos    int64
	index  int64
	ending bool
}

var mod int64

func pow(n int64, r int64) int64 {
	if r == 0 {
		return 1
	} else if r == 1 {
		return n % mod
	} else if r%2 == 0 {
		return pow((n*n)%mod, r/2)
	} else {
		return (n * pow((n*n)%mod, r/2)) % mod
	}
}
func main() {
	mod = 998244353
	var n, k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	endings := make([]Ending, 2*n)

	curr := 0
	for i := int64(0); i < n; i++ {
		var l, r int64

		fmt.Fscan(reader, &l, &r)
		endings[curr] = Ending{index: i, pos: l}
		curr++
		endings[curr] = Ending{index: i, pos: r, ending: true}
		curr++
	}
	sort.Slice(endings, func(i, j int) bool {
		return endings[i].pos < endings[j].pos || endings[i].pos == endings[j].pos && !endings[i].ending
	})
	all := map[int64]bool{}
	not_counted := map[int64]bool{}
	ans := int64(0)
	counts := map[int64]int64{}
	for i := int64(0); i < 2*n; i++ {

		if !endings[i].ending {
			all[endings[i].index] = true
			not_counted[endings[i].index] = true
		} else {
			if len(not_counted) > 0 {
				len64 := int64(len(all))

				if len64 >= k {
					counts[len64-k]++
					old := len64 - int64(len(not_counted))
					if old >= k {
						counts[old-k]--
					}
					not_counted = map[int64]bool{}

				}
			}
			delete(all, endings[i].index)
			if not_counted[endings[i].index] {
				delete(not_counted, endings[i].index)
			}
		}
	}
	if len(not_counted) > 0 {
		len64 := int64(len(all))
		if len64 >= k {
			counts[len64-k]++
			old := len64 - int64(len(not_counted))
			if old >= k {
				counts[old-k]--
			}
			not_counted = map[int64]bool{}
		}
	}
	fact, invs := make([]int64, n+1), make([]int64, n+1)

	fact[0] = 1
	for i := int64(1); i <= n; i++ {
		fact[i] = fact[i-1] * i
		fact[i] = fact[i] % mod
	}
	invs[n] = pow(fact[n], mod-2)
	for i := n - 1; i >= 0; i-- {
		invs[i] = invs[i+1] * (i + 1)
		invs[i] = invs[i] % mod
	}

	for key, v := range counts {
		if v != 0 {
			val := fact[key+k]
			val = val * invs[k]
			val = val % mod
			val = val * invs[key]
			val = val % mod
			if v > 0 {
				ans += v * val
				ans = ans % mod
			} else {
				ans += (mod + v) * val
				ans = ans % mod
			}
		}
	}
	write(f, ans, "\n")
}
