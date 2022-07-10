// 1367B. Even Array
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
		odds := 0
		misplaced := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if a[i]%2 == 1 {
				odds++
			}
			if i%2 == 1 && a[i]%2 == 0 {
				misplaced++
			}
		}
		if n/2 == odds {
			fmt.Println(misplaced)
		} else {
			fmt.Println(-1)
		}
	}
}
