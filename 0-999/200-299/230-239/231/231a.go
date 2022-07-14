// 231A - Team
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ans := 0
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)
		if a+b+c >= 2 {
			ans++
		}
	}
	fmt.Println(ans)

}
