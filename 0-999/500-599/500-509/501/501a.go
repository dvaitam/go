// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b, c, d int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &a, &b, &c, &d)
	vas := max(75*b, 250*b-b*d)
	mis := max(75*a, 250*a-a*c)
	if vas > mis {
		fmt.Println("Vasya")
	} else if mis > vas {
		fmt.Println("Misha")
	} else {
		fmt.Println("Tie")
	}

}
