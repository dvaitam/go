// 1352A. Sum of Round Numbers
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t, n int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		curr := 10
		l := []interface{}{}
		for n > 0 {
			r := n % curr
			if r != 0 {
				l = append(l, strconv.Itoa(r))
				n = n - r
			}
			curr = curr * 10
		}
		fmt.Println(len(l))
		fmt.Println(l...)
	}
}
