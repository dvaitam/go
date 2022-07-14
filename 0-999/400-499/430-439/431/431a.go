// 431A. Black Square
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := make([]int, 4)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 4; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var s string
	fmt.Fscan(reader, &s)
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += a[s[i]-'1']
	}
	fmt.Println(ans)
}
