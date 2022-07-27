// 275A. Lights Out
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := make([][]int, 3)
	b := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 3)
		b[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = 1
			fmt.Fscan(reader, &b[i][j])
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] += b[i][j]
			a[i][j] %= 2
			if i-1 >= 0 {
				a[i-1][j] += b[i][j]
				a[i-1][j] %= 2
			}
			if j-1 >= 0 {
				a[i][j-1] += b[i][j]
				a[i][j-1] %= 2
			}
			if i+1 < 3 {
				a[i+1][j] += b[i][j]
				a[i+1][j] %= 2
			}
			if j+1 < 3 {
				a[i][j+1] += b[i][j]
				a[i][j+1] %= 2
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}

}
