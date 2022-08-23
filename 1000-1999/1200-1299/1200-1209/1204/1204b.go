// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, l, r int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &l, &r)
	min := 0
	max := 0
	curr := 1
	for i := 0; i < n; i++ {
		if i < l {
			min += curr
			curr = curr * 2
		} else {
			min++
		}
	}
	curr = 1
	for i := 0; i < n; i++ {
		if i < r {
			max += curr
			curr = curr * 2
		} else {
			max += curr / 2
		}
	}

	fmt.Println(min, max)

}
