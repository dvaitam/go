// 709A. Juicer
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, b, d int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &b, &d)
	a := make([]int, n)
	sm := 0
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] <= b {
			sm += a[i]
			if sm > d {
				ans++
				sm = 0
			}
		}
	}
	fmt.Println(ans)

}
