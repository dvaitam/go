// 1206B. Make Product Equal One
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_cost(k int) int {
	if k > 1 {
		return k - 1
	} else if k < -1 {
		return (-1 * k) - 1
	} else {
		return 0
	}
}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	zeros := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] == 0 {
			zeros++
		}
	}
	cost := int64(0)
	for i := 0; i < n; i++ {
		icost := get_cost(a[i])
		if a[i] > 0 {
			a[i] -= icost
		} else if a[i] < 0 {
			a[i] += icost
		}
		cost += int64(icost)
	}
	product := 1
	for i := 0; i < n; i++ {
		product *= a[i]
	}
	if product == 0 {
		cost += int64(zeros)
	} else if product == -1 {
		cost += 2
	}
	fmt.Println(cost)
}
