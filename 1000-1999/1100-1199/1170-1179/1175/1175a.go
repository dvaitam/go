// 00
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
		var n, k int64
		fmt.Fscan(reader, &n, &k)
		ans := int64(0)
		for n > 0 {
			if n%k == 0 {
				n = n / k
				ans++
			} else {
				ans += (n % k)
				n = n - (n % k)
			}
		}
		fmt.Println(ans)
	}
}
