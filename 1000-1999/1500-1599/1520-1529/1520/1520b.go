// 1520B. Ordinary Numbers
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
		var n int
		fmt.Fscan(reader, &n)
		d := len(strconv.Itoa(n))
		ans := 0
		for i := 1; i < d; i++ {
			ans += 9
		}
		curr := 0
		cd := 0
		for cd < d {
			curr = curr * 10
			curr++
			cd++
		}
		start := curr
		for start <= n {
			ans++
			start += curr
		}
		fmt.Println(ans)

	}
}
