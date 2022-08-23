// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		c := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			if i == 0 {
				c[i] = a[i]
			} else {
				c[i] = c[i-1] + a[i]
			}
		}
		if c[n-1] == 0 {
			fmt.Println("NO")
		} else {
			for true {
				curr := n - 1
				for curr >= 0 && c[curr] != 0 {
					curr--
				}
				if curr < 0 {
					break
				}
				a = append(a[curr+1:], a[:curr+1]...)
				for i := 0; i < n; i++ {
					if i == 0 {
						c[i] = a[i]
					} else {
						c[i] = c[i-1] + a[i]
					}
				}
			}
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Print(a[i])
				if i == n-1 {
					fmt.Println()
				} else {
					fmt.Print(" ")
				}
			}
		}
	}
}
