// 00
package main

import (
	"bufio"
	"fmt"
	"os"
)

func write(f *bufio.Writer, a ...interface{}) {
	f.Write([]byte(fmt.Sprint(a...)))
}

type Number interface {
	int64 | float64 | int
}

func min[K Number](a K, b K) K {
	if a < b {
		return a
	}
	return b
}
func max[K Number](a K, b K) K {
	if a > b {
		return a
	}
	return b
}
func abs[K Number](a K) K {
	if a < 0 {
		return -a
	}
	return a
}

func get_coff(n int, a int, b int) (int, int) {
	// if a > b {
	// 	for k := 1; k <= b; k++ {
	// 		if (n - k*a) < 0 {
	// 			return -1, -1
	// 		}
	// 		if (n-k*a)%b == 0 {
	// 			return k, (n - k*a) / b
	// 		}
	// 	}
	// 	return -1, -1
	// } else {
	for k := 1; k <= a; k++ {
		if (n - k*b) < 0 {
			return -1, -1
		}
		if (n-k*b)%a == 0 {
			//	fmt.Println("yes ", k, a, b)
			return (n - k*b) / a, k
		}
	}
	return -1, -1
	//}
}
func is_valid(n int, a int, b int) bool {
	if a == 1 {
		return b == 1 || n%b == 1
	}
	curr := 1
	for curr <= n {
		if (n-curr)%b == 0 {
			return true
		}
		curr = curr * a
	}
	return false
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, a, b int
		fmt.Fscan(reader, &n, &a, &b)
		// if T == 100000 {
		// 	if t == 1779 {
		// 		write(f, n, a, b, "\n")
		// 	}
		// } else {
		if is_valid(n, a, b) {
			write(f, "Yes\n")
		} else {
			write(f, "No\n")
		}
		//}

	}
}
