// 427A. Police Recruits
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, curr int
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(tmp))
	sum := 0
	untreated := 0
	for i := 1; i <= n; i++ {
		if i == n {
			tmp, _ = reader.ReadString('\n')
		} else {
			tmp, _ = reader.ReadString(' ')
		}
		curr, _ = strconv.Atoi(strings.TrimSpace(tmp))
		if curr > 0 {
			sum += curr
		} else {
			if sum > 0 {
				sum--
			} else {
				untreated++
			}
		}
	}
	fmt.Println(untreated)

}
