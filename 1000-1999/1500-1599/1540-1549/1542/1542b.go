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
			fmt.Println("yes ", k, a, b)
			return (n - k*b) / a, k
		}
	}
	return -1, -1
	//}
}
func is_valid(n int, a int, b int) bool {
	if a == 1 {
		if n%b == 1 {
			return true
		} else {
			return false
		}
	}
	if b == 1 {
		return true
	}
	if n%b == 1 || n == a {
		return true
	}
	c1, c2 := get_coff(n, a, b)
	if c1 == -1 || c2 == -1 {
		return false
	}
	return is_valid(c1, a, b)
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
		if is_valid(n, a, b) {
			write(f, "Yes\n")
		} else {
			write(f, "No\n")
		}
	}
}
