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
		start_n := n
		for i := 0; i < len(primes); i++ {
			if n%primes[i] == 0 {
				ans[primes[i]]++
				n = n / primes[i]
				break
			}
		}
		if n == start_n {
			ans[n]++
			break
		}
	}
	return ans
}
func gcd(a int64, b int64) int64 {
	if b > a {
		return gcd(b, a)
	}
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}
func lcm(a int64, b int64) int64 {
	return a * b / gcd(a, b)
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

	var n, k int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int64, n)
	//	all_factors := map[int64]int{}
	ll := int64(1)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(reader, &a[i])
		ll = lcm(ll, gcd(k, a[i]))
	}

	if ll%k == 0 {
		write(f, "Yes\n")
	} else {
		write(f, "No\n")
	}
}
