// 1535A. Fair Playoff
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		a := make([]int, 4)
		m := make(map[int]int)
		for i := 0; i < 4; i++ {
			fmt.Fscan(reader, &a[i])
			m[a[i]] = i
		}
		sort.Ints(a)
		i, j := m[a[2]], m[a[3]]
		if i < 2 && j >= 2 || j < 2 && i >= 2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
