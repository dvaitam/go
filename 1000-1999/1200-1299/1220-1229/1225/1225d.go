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
			if primes[i]*primes[i] > n {
				ans[n]++
				n = 1
				break
			}
			if n%primes[i] == 0 {
				ans[primes[i]]++
				n = n / primes[i]
				break
			}
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
func max_power(x int64, cache map[int64]int) int {
	if x == 1 {
		return 1
	}
	if cache[x] > 0 {
		return cache[x]
	}
	ans := 0
	val := int64(1)
	for val <= 100000 {
		val *= x
		ans++
	}
	cache[x] = 2 * ans
	return 2 * ans
}
func pow(x int64, n int) int64 {
	val := int64(1)
	for n > 0 {
		val *= x
		n--
	}
	return val
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
	var n, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &k)
	a := make([]int64, n)
	m := map[int64]int64{}
	invm := map[int64]int64{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		//m[a[i]]++
		factors := prime_factors(a[i])
		val := int64(1)
		inv := int64(1)
		for key, v := range factors {
			v = v % k
			val *= pow(key, v)
			if v > 0 {
				inv *= pow(key, k-v)
			}
		}
		invm[val] = inv
		m[val]++
	}

	ans := int64(0)
	counted := map[int64]bool{}
	for key, v := range m {
		if counted[key] {
			continue
		}
		if key == invm[key] {
			ans += (v * (v - 1)) / 2
		} else {
			ans += v * m[invm[key]]
		}
		counted[invm[key]] = true
	}
	write(f, ans, "\n")

}
