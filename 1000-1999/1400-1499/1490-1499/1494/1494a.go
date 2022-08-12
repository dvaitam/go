// 1494A. ABC String
package main

import (
	"bufio"
	"fmt"
	"os"
)

func valid_sequence(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		A := "()"
		B := "()"
		C := "()"
		ok := false
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					res := []byte(s)
					m := make(map[byte]byte)
					m['A'] = A[i]
					m['B'] = B[j]
					m['C'] = C[k]
					for l := 0; l < len(s); l++ {
						res[l] = m[s[l]]
					}
					if valid_sequence(string(res)) {
						ok = true
					}
				}
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
