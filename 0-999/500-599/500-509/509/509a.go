// 509A. Maximum in Table
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &m)
	var a [10][10]int
	for i := 0; i < 10; i++ {
		a[0][i] = 1
		a[i][0] = 1
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			a[i][j] += a[i][j-1] + a[i-1][j]
		}
	}
	fmt.Println(a[m-1][m-1])

}
