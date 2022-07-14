// 231A - Team
package main

import "fmt"

func main() {
	var n, k, curr, kth int
	fmt.Scan(&n)
	fmt.Scan(&k)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&curr)
		if i < k-1 && curr > 0 {
			ans = i + 1
		} else if i == k-1 && curr > 0 {
			ans = i + 1
			kth = curr
		} else if i >= k && curr > 0 && curr == kth {
			ans = i + 1
		}
	}
	fmt.Println(ans)

}
