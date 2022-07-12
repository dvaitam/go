// 1703C. Cypher
package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_val(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'U' {
			ans--
		} else {
			ans++
		}
	}
	return ans
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
		}
		f := []interface{}{}
		for i := 0; i < n; i++ {
			var b int
			var s string
			fmt.Fscan(reader, &b, &s)
			f = append(f, (10+a[i]+get_val(s))%10)
		}
		fmt.Println(f...)

	}
}
