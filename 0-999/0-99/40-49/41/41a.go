// 41A. Translation
package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	i := 0
	j := len(s) - 1
	equal := true
	if len(s) == len(t) {
		for i < len(s) {
			if s[i] == t[j] {
				i++
				j--
			} else {
				equal = false
				break
			}
		}
	} else {
		equal = false
	}

	if equal {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
