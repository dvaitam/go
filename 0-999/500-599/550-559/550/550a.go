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
	var s string
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &s)
	fab, lab := -1, -1
	fba, lba := -1, -1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == 'A' && s[i+1] == 'B' {
				if fab == -1 {
					fab, lab = i, i
				} else {
					lab = i
				}
			}

			if s[i] == 'B' && s[i+1] == 'A' {
				if fba == -1 {
					fba, lba = i, i
				} else {
					lba = i
				}
			}
		}

	}
	if fab != -1 && fba != -1 && (abs(lab-fba) > 1 || abs(lba-fab) > 1) {
		write(f, "YES\n")
	} else {
		write(f, "NO\n")
	}

}
