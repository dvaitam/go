// 9A. Die Roll
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var m, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &m, &n)
	mx := max(m, n)
	if mx == 1 {
		fmt.Println("1/1")
	} else if mx == 2 {
		fmt.Println("5/6")
	} else if mx == 3 {
		fmt.Println("2/3")
	} else if mx == 4 {
		fmt.Println("1/2")
	} else if mx == 5 {
		fmt.Println("1/3")
	} else {
		fmt.Println("1/6")
	}

}
