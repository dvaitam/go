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
		var n, k int
		var s string
		fmt.Fscan(reader, &n, &k, &s)
		mk := 0
		i := 0
		j := n - 1
		for i < j {
			if s[i] != s[j] {
				break
			}
			i++
			j--
			mk++
		}
		if i == j+1 {
			mk--
		}
		//fmt.Println(mk, i, j)
		if k <= mk {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
