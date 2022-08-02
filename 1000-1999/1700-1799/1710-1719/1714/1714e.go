// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_group(k int) int {
	if k%10 == 5 {
		return k + 5
	}
	if k%10 == 0 {
		return k
	}
	if k == 1 || k%20 == 2 || k%20 == 4 || k%20 == 8 || k%20 == 16 {
		return 1
	}
	if k == 3 || k%20 == 6 || k%20 == 12 || k%20 == 14 || k%20 == 18 {
		return 3
	}
	return get_group(k + k%10)
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			a[i] = get_group(a[i])
		}
		ok := true
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
