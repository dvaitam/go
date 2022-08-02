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
		var n, curr int
		fmt.Fscan(reader, &n)
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			m[curr]++
		}
		mex_A := -1
		mex_B := -1
		for i := 0; i <= 100; i++ {
			if mex_A == -1 {
				if m[i] == 0 {
					mex_A = i
					mex_B = i
					break
				} else if m[i] == 1 {
					mex_A = i
				}
			} else {
				if m[i] == 0 {
					mex_B = i
					break
				}
			}
		}
		fmt.Println(mex_A + mex_B)
	}
}
