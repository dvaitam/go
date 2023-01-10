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
func pow(a int, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * pow(a, b-1)
	}
}
func get_letters(n int) string {
	ll := 1
	for l := 1; l <= 6; l++ {
		b := pow(26, l)
		if n > b {
			n = n - b
		} else {
			ll = l
			break
		}
	}

	parts := make([]int, ll)

	for l := ll; l > 0; l-- {
		b := pow(26, l-1)
		parts[ll-l] = n / b

		if n%b == 0 {
			parts[ll-l]--
		}
		if parts[ll-l] == -1 {
			parts[ll-l] = 25
		}
		n = n - (n/b)*b

	}

	ans := make([]byte, ll)
	i := 0

	//fmt.Println(parts)
	for i < ll {

		ans[i] = 'A' + byte(parts[i])
		i++
	}
	return string(ans)

}
func get_number(s string) int {
	ans := 0
	n := len(s)
	for l := 1; l < n; l++ {
		ans += pow(26, l)
	}
	for i := 0; i < n; i++ {
		if i == n-1 {
			ans += int(s[i] - 'A' + 1)
		} else {
			ans += int(s[i]-'A') * pow(26, n-i-1)
		}
	}
	return ans
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
		if s[0] == 'R' && len(s) > 1 && s[1] >= '0' && s[1] <= '9' {
			cindex := -1
			for i := 0; i < len(s); i++ {
				if s[i] == 'C' {
					cindex = i
					break
				}
			}
			if cindex == -1 {
				start := -1
				for i := 0; i < len(s); i++ {
					if s[i] >= '0' && s[i] <= '9' {
						start = i
						break
					}
				}
				row := s[start:]
				col := get_number(s[:start])
				write(f, "R"+row+"C"+strconv.Itoa(col), "\n")
			} else {
				row, _ := strconv.Atoi(s[1:cindex])
				col, _ := strconv.Atoi(s[cindex+1:])
				write(f, get_letters(col)+strconv.Itoa(row), "\n")
			}

		} else {
			start := -1
			for i := 0; i < len(s); i++ {
				if s[i] >= '0' && s[i] <= '9' {
					start = i
					break
				}
			}
			row := s[start:]
			col := get_number(s[:start])
			write(f, "R"+row+"C"+strconv.Itoa(col), "\n")
		}
	}
}
