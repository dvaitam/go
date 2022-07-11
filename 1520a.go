// 1520A. Do Not Be Distracted!
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
		var s string
		var n int
		fmt.Fscan(reader, &n, &s)
		so_far := make(map[byte]bool)
		distracted := false
		so_far[s[0]] = true
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] && so_far[s[i]] {
				distracted = true
				break
			}
			so_far[s[i]] = true
		}
		if distracted {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
