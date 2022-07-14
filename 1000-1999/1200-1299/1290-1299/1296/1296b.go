// 1296B. Food Buying
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_total(n int) int {
	if n < 10 {
		return n
	}
	return ((n / 10) * 10) + get_total((n/10)+n%10)
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s int
		fmt.Fscan(reader, &s)
		fmt.Println(get_total(s))
	}
}
