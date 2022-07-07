// 723A. The New Year: Meeting Friends
package main

import (
	"fmt"
)

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -1 * a
	}
}

func main() {
	var a [3]int
	fmt.Scan(&a[0], &a[1], &a[2])
	ans := 300
	for i := 1; i <= 100; i++ {
		d := 0
		d += abs(a[0] - i)
		d += abs(a[1] - i)
		d += abs(a[2] - i)
		if d < ans {
			ans = d
		}
	}

	fmt.Println(ans)

}
