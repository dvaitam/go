// 1650B. DIV + MOD
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var l, r, a int
		fmt.Fscan(reader, &l, &r, &a)
		if r%a != a-1 {
			left := r - r%a - 1
			if left >= l {
				fmt.Println(max(left/a+left%a, r/a+r%a))
			} else {
				fmt.Println(r/a + r%a)
			}
		} else {
			fmt.Println(r/a + r%a)
		}

	}
}
