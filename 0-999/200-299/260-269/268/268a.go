// 268A. Games
package main

import (
	"fmt"
)

func main() {
	var n, h, g int
	fmt.Scan(&n)
	a := make([]int, 30)

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&h, &g)
		a = append(a, h)
		m[g]++
	}
	ans := 0
	for _, v := range a {
		ans += m[v]
	}
	fmt.Println(ans)

}
