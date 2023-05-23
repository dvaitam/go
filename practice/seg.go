package main

import "fmt"

func lb(x int) int {
	return x & (-x)
}
func mod(bit []int, x int) {
	for i := x; i < len(bit); i += lb(i) {
		bit[i]++
	}
}
func query(bit []int, x int) int {
	r := 0
	for i := x; i > 0; i -= lb(i) {
		r += bit[i]
	}
	return r
}
func main() {
	bit := make([]int, 100)
	mod(bit, 26)
	for i := 0; i < len(bit); i++ {
		if bit[i] == 1 {
			fmt.Println(i)
		}
	}
	fmt.Println(bit)
}
