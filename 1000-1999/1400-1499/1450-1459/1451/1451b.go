// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve_method_one(s string, l int, r int) {
	n := len(s)
	sub := s[l-1 : r]
	//fmt.Println(sub)
	slen := len(sub)
	dp := make([][]int64, slen+1)
	for j := 0; j <= slen; j++ {
		dp[j] = make([]int64, n+1)
	}
	for j := 1; j <= slen; j++ {
		for k := 1; k <= n; k++ {
			if sub[j-1] == s[k-1] {
				if j == 1 {
					dp[j][k] = dp[j][k-1] + 1
				} else {
					dp[j][k] = dp[j][k-1] + dp[j-1][k-1]
				}
			} else {
				dp[j][k] = dp[j][k-1]
			}
			dp[j][k] = dp[j][k] % 1000001
		}
	}
	//fmt.Println(dp)
	count := int64(strings.Count(s, sub))
	if dp[slen][n] > count {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
func solve_method_two(s string, l int, r int) {
	l--
	r--
	ok := false
	for i := 0; i < l; i++ {
		if s[i] == s[l] {
			ok = true
		}
	}
	for i := r + 1; i < len(s); i++ {
		if s[i] == s[r] {
			ok = true
		}
	}
	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, q int
		fmt.Fscan(reader, &n, &q)
		var s string
		fmt.Fscan(reader, &s)
		for i := 0; i < q; i++ {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			solve_method_two(s, l, r)

		}
	}
}
