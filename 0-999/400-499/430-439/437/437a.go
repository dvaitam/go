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
func main() {
	var s1, s2, s3, s4 string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s1, &s2, &s3, &s4)
	n1, n2, n3, n4 := len(s1)-2, len(s2)-2, len(s3)-2, len(s4)-2
	// for i := 2; i < len(s1); i++ {
	// 	if s1[i] != '_' {
	// 		n1++
	// 	}
	// }
	// for i := 2; i < len(s2); i++ {
	// 	if s2[i] != '_' {
	// 		n2++
	// 	}
	// }
	// for i := 2; i < len(s3); i++ {
	// 	if s3[i] != '_' {
	// 		n3++
	// 	}
	// }
	// for i := 2; i < len(s4); i++ {
	// 	if s4[i] != '_' {
	// 		n4++
	// 	}
	// }
	A, B, C, D := 0, 0, 0, 0
	if 2*n1 <= n2 && 2*n1 <= n3 && 2*n1 <= n4 || n1 >= 2*n2 && n1 >= 2*n3 && n1 >= 2*n4 {
		A++
	}
	if 2*n2 <= n1 && 2*n2 <= n3 && 2*n2 <= n4 || n2 >= 2*n1 && n2 >= 2*n3 && n2 >= 2*n4 {
		B++
	}
	if 2*n3 <= n1 && 2*n3 <= n3 && 2*n3 <= n4 || n3 >= 2*n1 && n3 >= 2*n2 && n3 >= 2*n4 {
		C++
	}
	if 2*n4 <= n1 && 2*n4 <= n3 && 2*n4 <= n2 || n4 >= 2*n1 && n4 >= 2*n3 && n4 >= 2*n2 {
		D++
	}
	//write(f, A, B, D, "\n")
	if A+B+D+C == 1 {
		if A == 1 {
			write(f, "A\n")
		}
		if B == 1 {
			write(f, "B\n")
		}
		if D == 1 {
			write(f, "D\n")
		}
		if C == 1 {
			write(f, "C\n")
		}
	} else {
		write(f, "C\n")
	}

}
