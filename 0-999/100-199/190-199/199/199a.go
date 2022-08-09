// 199A. Hexadecimal's theorem
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
	a := make([]int, 2)
	a[0] = 0
	a[1] = 1
	ll := len(a)
	for a[ll-1] < n {
		a = append(a, a[ll-1]+a[ll-2])
		ll++
	}
	if n == 0 {
		fmt.Println(0, 0, 0)
	} else if n == 1 {
		fmt.Println(0, 0, 1)
	} else if n == 2 {
		fmt.Println(0, 1, 1)
	} else if n == 3 {
		fmt.Println(1, 1, 1)
	} else {
		fmt.Println(a[ll-5], a[ll-4], a[ll-2])
	}

}
