// 1714A. Everyone Loves to Sleep
package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, H, M, h, m int
		fmt.Fscan(reader, &n, &H, &M)
		minutes := 60*H + M
		min_greater := 60 * 24
		min_glob := 60 * 24
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &h, &m)
			mm := h*60 + m
			if mm >= minutes {
				min_greater = min(min_greater, mm)
			}
			min_glob = min(min_glob, mm)
		}
		if min_greater == 60*24 {
			ans := 24*60 - minutes + min_glob
			fmt.Println(ans/60, ans%60)
		} else {
			ans := min_greater - minutes
			fmt.Println(ans/60, ans%60)
		}

	}
}
