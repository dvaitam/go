// 732A. Buy a Shovel
package main

import (
	"fmt"
)

func main() {
	var k, r int
	fmt.Scan(&k, &r)
	ans := 1
	for k*ans%10 != 0 && (k*ans-r)%10 != 0 {
		ans++
	}
	fmt.Println(ans)
}
