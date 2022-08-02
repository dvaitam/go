// 1613A. Long Comparison
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var x1, p1, x2, p2 int
		fmt.Fscan(reader, &x1, &p1, &x2, &p2)
		digits1 := len(strconv.Itoa(x1)) + p1
		digits2 := len(strconv.Itoa(x2)) + p2
		if digits1 < digits2 {
			fmt.Println("<")
		} else if digits1 > digits2 {
			fmt.Println(">")
		} else {
			for len(strconv.Itoa(x1)) != len(strconv.Itoa(x2)) {
				if x1 < x2 {
					x1 = x1 * 10
				} else {
					x2 = x2 * 10
				}
			}
			if x1 < x2 {
				fmt.Println("<")
			} else if x1 > x2 {
				fmt.Println(">")
			} else {
				fmt.Println("=")
			}
		}

	}
}
