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
func pow(a int64, b int) int64 {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
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
func main() {
	primes = make([]int64, 1)
	primes[0] = 2
	for start := int64(3); start <= 31623; start++ {
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
	var l, r, x, y int64
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &l, &r, &x, &y)
	//write(f, prime_factors(6555736), "\n")
	if y % x != 0 {
		write(f, "0\n")
	}else{
		factors := prime_factors(y/x)
		pr := make([]int64, 0)
		for k := range factors {
			pr = append(pr, k)
		}
		ans := 0
		for i := 0; i < 1 <<(len(pr)); i++ {
			curr := int64(1)
			for j := 0; j < len(pr); j++ {
				if i & (1 << j) != 0 {
					curr = curr * pow(pr[j], factors[pr[j]])
				}
			}
			other := (y/x)/curr
			other = other*x
			curr = curr * x
			//write(f, curr, other, "\n")
			if curr >= l && curr <= r && other >= l && other <= r{
				ans++
			}
		}
		write(f, ans, "\n")

	}
}
