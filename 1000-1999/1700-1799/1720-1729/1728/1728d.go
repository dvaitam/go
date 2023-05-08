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

var ok bool
var cache [][]int

func check(s string, i int, j int, pick int) bool {

	if cache[i][j] > 0 {
		return cache[i][j] == 1
	}
	if j == i+1 {
		if s[i] != s[j] {
			return false
		} else {
			return true
		}
	} else {
		if (j-i)%2 == 1 {
			ans := check(s, i+1, j, i) && check(s, i, j-1, j)
			if ans {
				cache[i][j] = 1
			} else {
				cache[i][j] = 2
			}
			return ans
		} else {
			ans := true
			if s[pick] == s[i] && s[pick] == s[j] {
				ans = check(s, i+1, j, i) || check(s, i, j-1, j)
			} else if s[pick] == s[i] {
				ans = check(s, i+1, j, i)
			} else if s[pick] == s[j] {
				ans = check(s, i, j-1, j)
			} else {
				ans = false
			}
			if ans {
				cache[i][j] = 1
			} else {
				//fmt.Println(i, j, false)
				cache[i][j] = 2
			}
			return ans
		}
	}

}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		cache = make([][]int, len(s))
		for i := 0; i < len(s); i++ {
			cache[i] = make([]int, len(s))
		}
		ok = true

		if check(s, 0, len(s)-1, 0) {
			write(f, "Draw\n")
		} else {
			write(f, "Alice\n")
		}
	}
}
