// 1335A. Candies and Two Sisters
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t, n int
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	t, _ = strconv.Atoi(strings.TrimSpace(tmp))
	for i := 1; i <= t; i++ {
		tmp, _ = reader.ReadString('\n')
		n, _ = strconv.Atoi(strings.TrimSpace(tmp))
		if n%2 == 0 {
			fmt.Println(n/2 - 1)
		} else {
			fmt.Println(n / 2)
		}

	}

}
