// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, l, r int
		fmt.Fscan(reader, &n, &l, &r)
		ans := []interface{}{}
		ok := true
		for i := 1; i <= n; i++ {
			if l%i == 0 {
				ans = append(ans, l)
			} else {
				prev := l - (l % i)
				prev += i
				if prev <= r {
					ans = append(ans, prev)
				} else {
					ok = false
					break
				}
			}
		}
		if ok {
			fmt.Println("YES")
			fmt.Println(ans...)
		} else {
			fmt.Println("NO")
		}
	}
}
