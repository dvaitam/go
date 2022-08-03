// 1559B. Mocha and Red and Blue
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		var s string
		fmt.Fscan(reader, &s)
		a := []byte(s)
		i := 0
		for i < n && a[i] == '?' {
			i++
		}
		if i == n {
			blue := true
			for i > 0 {
				if blue {
					a[i-1] = 'B'
				} else {
					a[i-1] = 'R'
				}
				i--
				blue = !blue
			}
		}
		for i > 0 {
			if a[i] == 'B' {
				a[i-1] = 'R'
			} else {
				a[i-1] = 'B'
			}
			i--
		}
		for i < n {
			if a[i] == '?' {
				if a[i-1] == 'R' {
					a[i] = 'B'
				} else {
					a[i] = 'R'
				}
			}
			i++
		}
		fmt.Println(string(a))
	}
}
