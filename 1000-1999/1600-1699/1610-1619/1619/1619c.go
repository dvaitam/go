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
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var a, s string
		fmt.Fscan(reader, &a, &s)
		n := len(a)
		if len(s) < n {
			write(f, "-1\n")
			continue
		}
		ok := true
		ans := make([]rune, len(s))
		for i := 0; i < len(s); i++ {
			ans[i] = '0'
		}
		j := len(s) - 1
		p := len(s) - 1
		i := n - 1
		for ; ; i-- {

			ad := 0
			if i >= 0 {
				ad = int(a[i] - '0')
			}
			if j < 0 {
				break
			}
			sd := int(s[j] - '0')
			found := false
			for k := 0; k < 10; k++ {
				if ad+k == sd {
					ans[p] = rune(48 + k)
					found = true
					j--
					p--
					break
				}
			}
			if found {
				continue
			}
			if j-1 < 0 {
				ok = false
				break
			}
			sd, _ = strconv.Atoi(s[j-1 : j+1])
			for k := 0; k < 10; k++ {
				if ad+k == sd {
					ans[p] = rune(48 + k)
					found = true
					j -= 2
					p--
					break
				}
			}
			if !found {
				ok = false
				break
			}
		}
		//write(f, ans, string(ans), j, "\n")
		if ok {
			bi, _ := strconv.ParseInt(string(ans), 10, 64)
			if i >= 0 {
				write(f, "-1\n")
			} else {
				write(f, bi, "\n")
			}
		} else {
			write(f, "-1\n")
		}

	}
}
