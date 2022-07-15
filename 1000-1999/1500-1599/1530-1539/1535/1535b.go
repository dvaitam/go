// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a int, b int) int {
	if a > b {
		return gcd(b, a)
	}
	if b%a == 0 {
		return a
	} else {
		return gcd(b%a, a)
	}
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, curr int
		fmt.Fscan(reader, &n)
		odds := make([]int, 0)
		o := 0
		e := 0
		ans := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			if curr%2 == 0 {
				e++
			} else {
				o++
				odds = append(odds, curr)
			}
		}
		ans += o * e
		ans += (e * (e - 1)) / 2
		for i := 0; i < len(odds); i++ {
			for j := i + 1; j < len(odds); j++ {
				if gcd(odds[i], odds[j]) > 1 {
					ans++
				}
			}
		}
		fmt.Println(ans)
	}
}
