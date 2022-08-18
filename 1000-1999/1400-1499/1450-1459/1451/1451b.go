// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	}
}
