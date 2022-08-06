// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, curr int
		fmt.Fscan(reader, &n)
		p1, p2, p3 := -1, -1, -1
		m := make(map[int]int)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = i + 1
		}
		for i := 0; i < n; i++ {
			curr = a[i]
			if p1 == -1 {
				p1 = curr
			} else {
				if p2 == -1 {
					if curr <= p1 {
						p1 = curr
					} else {
						p2 = curr
					}
				} else {
					if curr >= p2 {
						p2 = curr
					} else {
						p3 = curr
						break
					}
				}
			}
		}
		if p3 != -1 {
			fmt.Println("YES")
			fmt.Println(m[p1], m[p2], m[p3])
			// fmt.Println(p1, p2, p3)
		} else {
			fmt.Println("NO")
		}

	}
}
