// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func write_int(f *bufio.Writer, a int) {
	f.Write([]byte(strconv.Itoa(a)))
}
func write_string(f *bufio.Writer, a string) {
	f.Write([]byte(a))
}

type Couple struct {
	a, b int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	total := 1
	for t := 1; t <= T; t++ {
		var A, B, n int
		fmt.Fscan(reader, &A, &B, &n)
		c := make([]Couple, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &c[i].a)
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &c[i].b)
		}

		sort.Slice(c, func(i, j int) bool { return c[i].a < c[j].a })

		i := 0
		for i < n {
			times := c[i].b / A
			if c[i].b%A != 0 {
				times++
			}

			// j := 0
			// for j < times {
			// 	B = B - c[i].a
			// 	j++
			// 	if B <= 0 {
			// 		break
			// 	}
			// }
			// if j != times {
			// 	break
			// }

			can_fight := B / c[i].a
			if B%c[i].a != 0 {
				can_fight++
			}
			if can_fight < times {
				break
			} else {
				B = B - times*c[i].a
			}

			// if B-(times-1)*c[i].a > 0 {
			// 	B = B - c[i].a*times
			// } else {
			// 	break
			// }

			i++
			if B <= 0 {
				break
			}
		}
		if i == n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		total++
	}
}
