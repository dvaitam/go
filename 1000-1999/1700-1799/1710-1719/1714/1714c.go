// 1714C. Minimum Varied Number
package main

import (
	"bufio"
	"fmt"
	"os"
)

func print_ans(curr int, max int) {
	if curr <= max {
		fmt.Print(curr)
	} else {
		print_ans(curr-max, max-1)
		fmt.Print(max)
	}
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s int
		fmt.Fscan(reader, &s)
		print_ans(s, 9)
		fmt.Println()
	}
}
