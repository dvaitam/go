// 1703A. YES or YES?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		if strings.ToLower(s) == "yes" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
