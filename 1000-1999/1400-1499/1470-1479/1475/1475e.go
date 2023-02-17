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

var primes []int64

func prime_factors(n int64) map[int64]int {
	ans := map[int64]int{}
	for n > 1 {
		for i := 0; i < len(primes); i++ {
			if n%primes[i] == 0 {
				ans[primes[i]]++
				n = n / primes[i]
				break
			}
		}
	}
	return ans
}
func ncr(n int64, r int64) int64 {
	if n < r {
		return 0
	}
	if n == r {
		return 1
	}
	if r == 0 {
		return 1
	}
	nu := map[int64]int{}
	tmp := n

	for d := r; d > 0; d-- {
		m := prime_factors(tmp)
		for k, v := range m {
			nu[k] += v
		}
		tmp--
	}
	for d := r; d > 0; d-- {
		m := prime_factors(d)
		for k, v := range m {
			nu[k] -= v
		}
	}
	ans := int64(1)

	for k, v := range nu {
		for j := 0; j < v; j++ {
			ans = ans * k
			ans = ans % 1000000007
		}
	}

	return ans
}

func main() {
	primes = make([]int64, 1)
	primes[0] = 2
	for start := int64(3); start <= 1000; start++ {
		ok := true
		for i := 0; i < len(primes); i++ {
			if start%primes[i] == 0 {
				ok = false
				break
			}
		}
		if ok {
			primes = append(primes, start)
		}
	}
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int64
		fmt.Fscan(reader, &n, &k)
		a := make([]int64, n)
		m := map[int64]int64{}
		for i := int64(0); i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]]++
		}
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		s := int64(0)
		include := int64(0)
		for i := n - 1; i >= n-k; i-- {
			s += a[i]
			if a[i] == a[n-k] {
				include++
			}
		}
		write(f, ncr(m[a[n-k]], include), "\n")
	}
}
