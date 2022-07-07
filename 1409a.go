// 1409A. Yet Another Two Integers Problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T, a, b int
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	T, _ = strconv.Atoi(strings.TrimSpace(tmp))
	for t := 0; t < T; t++ {
		tmp, _ = reader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(tmp), " ")
		a, _ = strconv.Atoi(parts[0])
		b, _ = strconv.Atoi(parts[1])
		d := a - b
		if d < 0 {
			d = b - a
		}
		ans := 0
		for i := 10; i > 0; i-- {
			ans += (d / i)
			d = d - (d/i)*i
		}
		fmt.Println(ans)
	}

}
