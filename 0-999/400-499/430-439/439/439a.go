// 439A. Devu, the Singer and Churu, the Joker
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var T int
	reader := bufio.NewReader(os.Stdin)
	// fmt.Fscan(reader, &T)
	// for t := 1; t <= T; t++ {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		sum += a[i]
		if i > 0 {
			sum += 10
		}
	}
	if sum <= d {
		fmt.Println(2*n - 2 + (d-sum)/5)
	} else {
		fmt.Println(-1)
	}
	//}
}
