// 750A. New Year and Hurry
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	k = 240 - k
	sq := math.Sqrt(float64(25 + 40*k))
	ans := int((sq - 5) / 10)
	if ans <= n {
		fmt.Println(ans)
	} else {
		fmt.Println(n)
	}

}
