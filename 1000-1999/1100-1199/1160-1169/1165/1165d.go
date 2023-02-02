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

var primes map[int]bool
var smallest map[int]int

func is_prime(n int) bool {
	if primes[n] {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	primes[n] = true
	return true
}
func prime_factors(n int, counter map[int]int) {
	if n == 1 {
		return
	}
	small := smallest[n]
	if small == 0 {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				//if is_prime(i) {
				small = i
				break
				//}
			}
		}
	}
	if small == 0 {
		small = n
	}
	smallest[n] = small

	counter[small]++
	prime_factors(n/small, counter)
}
func pow(a int, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
}
func main() {
	primes = make(map[int]bool)
	smallest = make(map[int]int)

	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		d := make([]int, n)
		md := map[int]bool{}
		smallest_factor := 10000000
		largest_factor := 1
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &d[i])
			smallest_factor = min(smallest_factor, d[i])
			largest_factor = max(largest_factor, d[i])
			md[d[i]] = true
		}
		max_map := map[int]int{}
		ok := true

		for i := 0; i < n; i++ {
			counter := make(map[int]int)
			prime_factors(d[i], counter)
			//	write(f, d[i], counter, "\n")
			for k, v := range counter {
				if !md[k] {
					ok = false
					break
				}
				max_map[k] = max(max_map[k], v)
			}
			if !ok {
				break
			}
		}
		//write(f, max_map, ok, "\n")
		ans := 1
		if ok {
			total_factors := 1
			for k, v := range max_map {
				total_factors *= (v + 1)
				ans *= pow(k, v)
			}
			if ans == largest_factor {
				max_map[smallest_factor]++
				total_factors = 1
				ans = 1
				for k, v := range max_map {
					total_factors *= (v + 1)
					ans *= pow(k, v)
				}

			}
			if total_factors != n+2 {
				ok = false
			}

		}
		if ok {
			write(f, ans, "\n")
		} else {
			write(f, "-1\n")
		}
	}
}
