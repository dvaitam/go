// 00
package main

import (
	"bufio"
	"fmt"
	"os"
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
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func main() {

	//reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	a, b := int64(1), int64(2)
	for t := 1; t <= 50; t++ {
		var l1, l2 int64
		fmt.Println("? ", a, b)
		fmt.Scan(&l1)
		fmt.Println("? ", b, a)
		fmt.Scan(&l2)
		if l1 == -1 || l2 == -1 {
			a, b = int64(1), int64(2)
			continue
		}
		if l1 != l2 {
			fmt.Println("! ", l1+l2)
			break
		}
		a++
		b++
	}
}
