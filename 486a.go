// 486A. Calculating Function
package main

import (
	"fmt"
)

func main() {
	var n, ans int64
	fmt.Scan(&n)
	if n%2 == 0 {
		ans = (n/2)*(n/2+1) - (n/2)*(n/2)
	} else {
		ans = (n/2)*(n/2+1) - (n/2+1)*(n/2+1)
	}
	fmt.Println(ans)

}
