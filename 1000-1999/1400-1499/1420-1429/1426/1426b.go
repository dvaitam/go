// 1426B. Symmetric Matrix
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
		var n, m int
		fmt.Fscan(reader, &n, &m)
		ok := false
		for i := 0; i < n; i++ {
			var a, b, c, d int
			fmt.Fscan(reader, &a, &b, &c, &d)
			if b == c {
				ok = true
			}
		}
		if m%2 == 1 || ok == false {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
