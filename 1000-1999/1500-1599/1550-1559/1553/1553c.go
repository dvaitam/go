// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func get_binary(a int) string {
	s := strconv.FormatInt(int64(a), 2)
	ans := make([]byte, 10)
	start := 10 - len(s)
	j := 0
	for i := 0; i < 10; i++ {
		if i < start {
			ans[i] = '0'
		} else {
			ans[i] = s[j]
			j++
		}
	}
	return string(ans)
}
func get_min(s string) int {
	first := 0
	second := 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			if s[i] == '1' {
				first++
			}
		} else {
			if s[i] == '1' {
				second++
			}
		}
		rem := 10 - (i + 1)
		if first+(rem/2) < second {
			return i + 1
		}
		ss := rem / 2
		if rem%2 == 1 {
			ss++
		}
		if second+ss < first {
			return i + 1
		}
	}
	return 10
}
func equal(s1 string, s2 string) bool {
	for i := 0; i < 10; i++ {
		if s1[i] == '?' || s1[i] == s2[i] {
			continue
		} else {
			return false
		}
	}
	return true
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	cache := map[string]int{}
	for i := 0; i < 1<<10; i++ {
		s := get_binary(i)
		cache[s] = get_min(s)
	}
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		ans := 10
		for i := 0; i < 1<<10; i++ {
			s2 := get_binary(i)
			if equal(s, s2) {
				ans = min(ans, cache[s2])
			}
		}
		write(f, ans, "\n")
	}
}
