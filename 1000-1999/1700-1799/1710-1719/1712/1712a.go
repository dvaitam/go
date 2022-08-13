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
		fmt.Fscan(reader, &n, &k)
		a := make([]int, n)
		ans := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i < k && a[i] > k {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
