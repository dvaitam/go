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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		var s1, s2 string
		fmt.Fscan(reader, &n, &s1, &s2)

		ans := 0
		found_zero := 0
		found_one := 0
		for i := 0; i < n; i++ {
			if s1[i] == '0' && s2[i] == '1' || s1[i] == '1' && s2[i] == '0' {
				ans += found_zero + 2
				found_one = 0
				found_zero = 0
			} else if found_zero > 0 && s1[i] == '1' && s2[i] == '1' {
				ans += found_zero + 1
				found_zero = 0
				found_one = 0
			} else if found_one > 0 && s1[i] == '0' && s2[i] == '0' {
				ans += 2
				found_zero = 0
				found_one = 0
			} else if s1[i] == '0' && s2[i] == '0' {
				found_zero++
			} else if s1[i] == '1' && s2[i] == '1' {
				found_one++
			}
		}
		ans += found_zero
		write(f, ans, "\n")
	}
}
