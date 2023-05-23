// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}
func bits(n int) int {
	ans := 0
	for n > 0 {
		if n&1 == 1 {
			ans++
		}
		n = n / 2
	}
	return ans

}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var cnt int
		fmt.Fscan(reader, &cnt)
		ans := 0
		tmp := 1
		//ans += tmp
		last_count := cnt
		i := 1
		surplus := 0
		for cnt > 0 {
			fmt.Println("- ", tmp)
			fmt.Fscan(reader, &cnt)
			ans += tmp

			if cnt == last_count-1 {

			} else {
				surplus += tmp
			}
			tmp = 1 << i
			if cnt == bits(surplus) {
				tmp = surplus
				surplus = 0
			}
			last_count = cnt
			i++

		}
		fmt.Println("! ", ans)
	}
}
