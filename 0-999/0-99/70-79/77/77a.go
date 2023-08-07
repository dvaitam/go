// 00
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	l := make([][]bool, 7)
	for i := 0; i < 7; i++ {
		l[i] = make([]bool, 7)
	}
	m := map[string]int{"Anka": 0, "Chapay": 1, "Cleo": 2, "Troll": 3, "Dracul": 4, "Snowy": 5, "Hexadecimal": 6}
	for i := 0; i < n; i++ {
		var s1, s2, s3 string
		fmt.Fscan(reader, &s1, &s2, &s3)
		l[m[s1]][m[s3]] = true
	}
	a := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Ints(a)
	ml := make([][]int, 128)
	em := make([][]int, 128)
	for i := 0; i < 128; i++ {
		ml[i] = make([]int, 128)
		em[i] = make([]int, 128)
		for j := 0; j < 128; j++ {
			em[i][j] = -1
		}
	}
	ans_0 := 2000000000
	for i := 1; i < 128; i++ {
		for j := 1; j < 128; j++ {
			if i&j != 0 {
				continue
			}
			k := (1 << 7) - 1
			k = k ^ (i | j)
			if k == 0 {
				continue
			}
			like := 0
			for p := 0; p < 7; p++ {
				for q := 0; q < 7; q++ {
					if p == q {
						continue
					}
					if l[p][q] && i&(1<<p) != 0 && i&(1<<q) != 0 {
						like++
					}
					if l[p][q] && j&(1<<p) != 0 && j&(1<<q) != 0 {
						like++
					}
					if l[p][q] && k&(1<<p) != 0 && k&(1<<q) != 0 {
						like++
					}
				}
			}
			ml[i][j] = like
			is := 0
			js := 0
			ks := 0
			for p := 0; p < 7; p++ {
				if i&(1<<p) != 0 {
					is++
				}
				if j&(1<<p) != 0 {
					js++
				}
				if k&(1<<p) != 0 {
					ks++
				}
			}
			ls := []int{is, js, ks}
			//write(f, ls, "\n")
			sort.Ints(ls)
			m1, m2, m3 := a[0]/ls[0], a[1]/ls[1], a[2]/ls[2]
			em[i][j] = max(max(m1, m2), m3) - min(min(m1, m2), m3)
			ans_0 = min(ans_0, em[i][j])
		}
	}
	ans_1 := 0
	for i := 0; i < 127; i++ {
		for j := 0; j < 127; j++ {
			if em[i][j] == ans_0 {
				ans_1 = max(ans_1, ml[i][j])
			}
		}
	}
	write(f, ans_0, ans_1, "\n")
}
