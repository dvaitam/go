// 1560A. Dislike of Threes
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T, k int
	l := []interface{}{}
	for i := 1; len(l) < 1000; i++ {
		if i%3 == 0 || i%10 == 3 {
			continue
		} else {
			l = append(l, i)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		fmt.Fscan(reader, &k)
		fmt.Println(l[k-1])
	}

}
