// 00
package main

import (
	"bufio"
	"fmt"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func main() {
	var T int
	// reader := bufio.NewReader(os.Stdin)
	// f := bufio.NewWriter(os.Stdout)
	// defer f.Flush()
	fmt.Scan(&T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Scan(&n, &m)
		var m1, m2 int
		fmt.Println("? ", 1, 1)
		fmt.Scan(&m1)
		fmt.Println("? ", n, 1)
		fmt.Scan(&m2)
		if m1+m2 > n-1 {
			if m1 < m2 {
				fmt.Println("! ", n-m2, m1+1)
			} else if m2 < m1 {
				fmt.Println("! ", m1+1, m2+1)
			} else {
				fmt.Println(" ?", 1, m1+1)
				var x int
				fmt.Scan(&x)
				fmt.Println("! ", x+1, m1+1)
			}
		} else {
			fmt.Println("? ", m1+1, 1)
			var x int
			fmt.Scan(&x)
			fmt.Println("! ", m1+1, x+1)

		}
	}
}
