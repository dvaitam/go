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
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	var s1, s2, s3 string
	fmt.Fscan(reader, &s1, &s2, &s3)
	m := len(s1)
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	m3 := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]]++
		m2[s2[i]]++
		m3[s3[i]]++
	}
	m1x, m2x, m3x := 0, 0, 0
	for _, v := range m1 {
		m1x = max(m1x, v)
	}
	for _, v := range m2 {
		m2x = max(m2x, v)
	}
	for _, v := range m3 {
		m3x = max(m3x, v)
	}
	//write(f, m1x, m2x, m3x, m, "\n")
	if m1x+n > m {
		if m1x == m && n == 1 {
			m1x = m - 1
		} else {
			m1x = m
		}
	} else {
		m1x = m1x + n
	}
	if m2x+n > m {
		if m2x == m && n == 1 {
			m2x = m - 1
		} else {
			m2x = m
		}
	} else {
		m2x = m2x + n
	}
	if m3x+n > m {
		if m3x == m && n == 1 {
			m3x = m - 1
		} else {
			m3x = m
		}
	} else {
		m3x = m3x + n
	}
	mmx := max(m1x, max(m2x, m3x))
	counts := map[int]int{}
	counts[m1x]++
	counts[m2x]++
	counts[m3x]++
	if counts[mmx] > 1 {
		write(f, "Draw\n")
	} else if mmx == m1x {
		write(f, "Kuro\n")
	} else if mmx == m2x {
		write(f, "Shiro\n")
	} else {
		write(f, "Katie\n")
	}

}
