// 1525B. Permutation Sort
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
		sorted := true
		for i := 1; i < n; i++ {
			if a[i] < a[i-1] {
				sorted = false
				break
			}
		}
		if sorted {
			fmt.Println(0)
		} else {
			if a[0] == n && a[n-1] == 1 {
				fmt.Println(3)
			} else if a[0] == 1 || a[n-1] == n {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		}
	}
}
