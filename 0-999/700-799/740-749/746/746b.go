// 746B. Decoding
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var T int
	reader := bufio.NewReader(os.Stdin)
	//fmt.Fscan(reader, &T)
	//for t := 1; t <= T; t++ {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	ans := make([]byte, n)
	left := (n / 2) - 1
	right := left + 1
	start := 0

	if n%2 == 1 {
		ans[left+1] = s[0]
		right = left + 2
		start++
	}

	for i := start; i < n; i++ {
		if i%2 != n%2 {
			ans[right] = s[i]
			right++
		} else {
			ans[left] = s[i]
			left--
		}
	}
	fmt.Println(string(ans))
	//}
}
