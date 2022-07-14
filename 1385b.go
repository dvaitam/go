// 1385B. Restore the Permutation by Merger
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
		var n, curr int
		fmt.Fscan(reader, &n)
		a := make([]bool, n)
		l := []interface{}{}
		for i := 0; i < 2*n; i++ {
			fmt.Fscan(reader, &curr)
			if !a[curr-1] {
				l = append(l, curr)
				a[curr-1] = true
			}
		}
		fmt.Println(l...)
	}
}
