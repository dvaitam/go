// 1486A. Shifting Stacks
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
		curr := int64(0)
		sum := int64(0)
		ok := true
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &curr)
			sum += curr
			if sum < int64((i+1)*i)/2 {
				ok = false
			}

		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
