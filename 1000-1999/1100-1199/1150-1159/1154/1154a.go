// 1154A. Restoring Three Numbers
package main

import (
	"fmt"
)

func main() {
	a := [4]int{}
	max := 0
	max_index := 0
	for i := 0; i < 4; i++ {
		fmt.Scan(&a[i])
		if a[i] > max {
			max = a[i]
			max_index = i
		}
	}
	l := []interface{}{}
	for i := 0; i < 4; i++ {
		if i != max_index {
			l = append(l, max-a[i])
		}
	}
	fmt.Println(l...)

}
