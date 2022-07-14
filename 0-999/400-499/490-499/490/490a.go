// 490A. Team Olympiad
package main

import (
	"fmt"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, curr int
	fmt.Scan(&n)
	l1 := []interface{}{}
	l2 := []interface{}{}
	l3 := []interface{}{}
	for i := 1; i <= n; i++ {
		fmt.Scan(&curr)
		if curr == 1 {
			l1 = append(l1, i)
		} else if curr == 2 {
			l2 = append(l2, i)
		} else {
			l3 = append(l3, i)
		}
	}
	n1 := len(l1)
	n2 := len(l2)
	n3 := len(l3)
	ans := min(n1, min(n2, n3))
	fmt.Println(ans)
	for i := 0; i < ans; i++ {
		fmt.Println(l1[i], l2[i], l3[i])
	}
}
