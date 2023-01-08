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
func get_val(n int64, k int64) int64 {
	mi := n / k
	if n%k != 0 {
		mi++
	}
	mx := n
	ans := int64(0)
	for e := mi; e <= mx; e++ {
		for i := int64(0); i <= e; i++ {
			curr := e + k*i
			if curr <= n {
				rem := n - curr
				inter := ncr(e, i) * ncr(e+rem-1, rem)
				inter = inter % 1000000007
				if i%2 == 0 {
					ans += inter
				} else {
					ans -= inter
				}
				if ans < 0 {
					ans += 1000000007
				}
				ans = ans % 1000000007
			}

		}
	}
	return ans
}
func main() {
	primes = make([]int64, 1)
	primes[0] = 2
	for start := int64(3); start <= 100; start++ {
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
	var n, k, d int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k, &d)
	ans := get_val(n, k)
	if d > 1 {
		ans -= get_val(n, d-1)
	}
	if ans < 0 {
		ans += 1000000007
	}
	ans = ans % 1000000007
	write(f, ans, "\n")

}
