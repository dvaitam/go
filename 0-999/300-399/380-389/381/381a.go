// 381A. Sereja and Dima
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}
	first := true
	sereja := 0
	dime := 0
	f := 0
	e := n - 1
	for f <= e {
		if first {
			if l[f] > l[e] {
				sereja += l[f]
				f++
			} else {
				sereja += l[e]
				e--
			}
		} else {
			if l[f] > l[e] {
				dime += l[f]
				f++
			} else {
				dime += l[e]
				e--
			}
		}
		first = !first
	}
	fmt.Println(sereja, dime)
}
