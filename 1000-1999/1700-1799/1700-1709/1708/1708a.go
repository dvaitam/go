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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		ok := true
		for i := 1; i < n; i++ {
			if a[i]%a[0] != 0 {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
