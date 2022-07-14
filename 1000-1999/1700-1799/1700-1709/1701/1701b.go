// 1701B. Permutation
package main

import (
	"fmt"
)

func main() {
	var T, n int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		fmt.Scan(&n)
		fmt.Println(2)
		m := make(map[int]bool)
		l := []interface{}{}
		for i := 1; i <= n/2; i++ {
			if !m[i] && !m[i*2] {
				curr := i
				for curr <= n {
					m[curr] = true
					l = append(l, curr)
					curr = curr * 2
				}
			}
		}
		for i := 1; i <= n; i++ {
			if !m[i] {
				l = append(l, i)
			}
		}
		fmt.Println(l...)
	}
}
