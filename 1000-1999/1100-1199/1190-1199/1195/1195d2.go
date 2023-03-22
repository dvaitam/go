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

var d int64

func pow(a int64, b int64) int64 {
	if b == 0 {
		return 1
	} else {
		return (a * pow(a, b-1)) % d
	}
}

func get_val(s string, l int64) int64 {
	ans := int64(0)
	start := int64(0)
	for i := len(s) - 1; i >= 0; i-- {
		ans += int64(s[i]-'0') * pow(10, start)

		if int64(len(s)-1-i) < l {
			start += 2
		} else {
			start++
		}
		ans = ans % d
	}
	return ans
}
func get_val_other(s string, l int64) int64 {
	ans := int64(0)
	start := int64(1)
	for i := len(s) - 1; i >= 0; i-- {
		ans += int64(s[i]-'0') * pow(10, start)

		if int64(len(s)-i) < l {
			start += 2
		} else {
			start++
		}
		ans = ans % d
	}
	return ans
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	a := make([]string, n)
	d = int64(998244353)
	counts := map[int]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		counts[len(a[i])]++
	}
	ans := int64(0)
	for i := 0; i < n; i++ {
		for k, v := range counts {
			val := get_val(a[i], int64(k))
			val = val % d
			ans += (val * int64(v)) % d
			ans = ans % d
			val = get_val_other(a[i], int64(k))
			val = val % d
			ans += (val * int64(v)) % d
			ans = ans % d
		}
	}
	write(f, ans, "\n")
}
