// 1472B. Fair Division
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
		var n, curr int
		fmt.Fscan(reader, &n)
		ones := 0
		twos := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			if curr == 1 {
				ones++
			} else {
				twos++
			}
		}
		if ones%2 == 1 {
			fmt.Println("NO")
		} else if ones == 0 && twos%2 == 1 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
