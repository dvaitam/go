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
	mins := map[int64]int64{0: 0, 4: 1, 6: 1, 8: 2, 10: 2, 12: 2, 14: 3}
	maxes := map[int64]int64{0: 0, 4: 1, 6: 1, 8: 2, 10: 2, 12: 3, 14: 3}
	for t := 1; t <= T; t++ {
		var n int64
		fmt.Fscan(reader, &n)
		if n%2 == 1 || n == 2 {
			fmt.Println(-1)
		} else {
			if n%12 == 2 {
				fmt.Println((n-14)/6+mins[14], (n-14)/4+maxes[14])
			} else {
				rem := n - (n % 12)
				fmt.Println(rem/6+(mins[n%12]), rem/4+(maxes[n%12]))
			}
		}

	}
}
