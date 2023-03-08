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

func get_smallest(n int64) int64 {
	for i := 0; i < len(primes); i++ {
		if primes[i]*primes[i] > n {
			return n
		}
		if n%primes[i] == 0 {
			return primes[i]
		}
	}
	return n
}
func main() {
	prime_map := map[int64]int64{}
	primes = make([]int64, 1)
	primes[0] = 2

	prime_map[2] = 1
	counter := int64(2)
	for start := int64(3); start <= 199999; start++ {
		ok := true
		for i := 0; i < len(primes); i++ {
			if start%primes[i] == 0 {
				ok = false
				break
			}
		}
		if ok {
			primes = append(primes, start)
			prime_map[start] = counter
			counter++
		}
	}
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]int64, 2*n)
	counts := map[int64]int64{}
	for i := 0; i < 2*n; i++ {
		fmt.Fscan(reader, &a[i])
		counts[a[i]]++
	}
	//fmt.Println(get_smallest(2750130))
	//first := a[0]
	sort.Slice(a, func(i, j int) bool { return a[i] >= a[j] })
	ans := make([]int64, 0)
	rem := make([]int64, 0)
	removed := 0
	for i := 0; i < 2*n; i++ {
		if counts[a[i]] > 0 {
			small := get_smallest(a[i])
			if small != a[i] {
				factor := a[i] / small
				counts[factor]--
				ans = append(ans, a[i])
				counts[a[i]]--
			}
		}
	}
	for k, v := range counts {
		if v == 0 {
			continue
		}
		for i := int64(0); i < v; i++ {
			rem = append(rem, k)
		}
	}
	sort.Slice(rem, func(i, j int) bool { return rem[i] < rem[j] })

	not_handled := make([]int64, 0)
	for i := 0; i < len(rem); i++ {
		if counts[rem[i]] > 0 {
			if rem[i]-1 < int64(len(primes)) && counts[primes[rem[i]-1]] > 0 { //&& counts[primes[rem[i]-1]] > 0
				counts[rem[i]]--
				counts[primes[rem[i]-1]]--
				removed += 2
				ans = append(ans, rem[i])
			}
		}
	}
	for k, v := range counts {
		if v == 0 {
			continue
		}
		for i := int64(0); i < v; i++ {
			not_handled = append(not_handled, k)
		}
	}
	sort.Slice(not_handled, func(i, j int) bool { return not_handled[i] < not_handled[j] })

	for i := 0; i < len(not_handled)/2; i++ {
		ans = append(ans, not_handled[i])
	}

	//fmt.Println(rem, not_handled)
	// if first == 1648739 {
	// 	fmt.Println(len(ans), len(not_handled), len(rem), removed)
	// 	fmt.Println(not_handled[:len(not_handled)/2])
	// 	fmt.Println(not_handled[len(not_handled)/2:])
	// }

	sort.Slice(ans, func(i, j int) bool { return ans[i] > ans[j] })
	for i := 0; i < n; i++ {
		write(f, ans[i], " ")
	}
	write(f, "\n")
}
