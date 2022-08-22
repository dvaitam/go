// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Tuple struct {
	val, index int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	curr := 0
	a := make([]Tuple, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &curr)
		a[i] = Tuple{val: curr, index: i + 1}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].val > a[j].val })
	ans := 0
	x := 0
	for i := 0; i < n; i++ {
		ans += a[i].val*x + 1
		x++
	}
	fmt.Println(ans)
	for i := 0; i < n; i++ {
		fmt.Print(a[i].index, " ")
	}
	fmt.Println()

}
