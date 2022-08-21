// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	odds_count := 0
	evens_count := 0
	curr := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &curr)
		if curr%2 == 1 {
			odds_count++
		} else {
			evens_count++
		}
	}
	if odds_count%2 == 1 {
		fmt.Println(odds_count)
	} else {
		fmt.Println(evens_count)
	}

}
