// 1702D. Not a Cheap String
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		// s = strings.TrimSpace(s)
		var p int
		fmt.Fscan(reader, &p)
		total := 0
		m := make(map[int]int)
		for i := 0; i < len(s); i++ {
			val := int(s[i] - 'a' + 1)
			m[val]++
			total += val
		}

		if total > p {
			for i := 26; i > 0; i-- {
				if m[i] > 0 {
					if m[i]*i <= total-p {
						total -= m[i] * i
						m[i] = 0
					} else {
						rm := (total - p) / i
						if (total-p)%i != 0 {
							rm++
						}
						m[i] -= rm
						total -= rm * i
					}
					if total <= p {
						break
					}
				}
			}
		}

		l := make([]byte, 0)
		for i := 0; i < len(s); i++ {
			val := int(s[i] - 'a' + 1)
			if m[val] > 0 {
				l = append(l, s[i])
				m[val]--
			}
		}
		fmt.Println(string(l))
	}
}
