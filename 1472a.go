// 1472A. Cards for Friends
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_count(w int, h int) int {
	if w%2 == 0 {
		return 2 * get_count(w/2, h)
	} else if h%2 == 0 {
		return 2 * get_count(w, h/2)
	} else {
		return 1
	}
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var w, h, n int
		fmt.Fscan(reader, &w, &h, &n)
		if get_count(w, h) < n {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
