// 1705B. Mark the Dust Sweeper
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
		start_count := false
		count := 0
		sm := int64(0)
		for i := 0; i < n-1; i++ {
			if a[i] != 0 {
				start_count = true
				sm += int64(a[i])
			} else if start_count {
				count++
			}
		}
		fmt.Println(int64(count) + sm)

	}
}
