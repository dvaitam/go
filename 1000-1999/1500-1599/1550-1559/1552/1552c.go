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

type Chord struct {
	p1, p2 int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		chords := make([]Chord, k)
		done := map[int]bool{}
		for i := 0; i < k; i++ {
			var p1, p2 int
			fmt.Fscan(reader, &p1, &p2)
			if p1 < p2 {
				chords[i] = Chord{p1, p2}
			} else {
				chords[i] = Chord{p2, p1}

			}
			done[p1] = true
			done[p2] = true
		}
		rem := make([]int, 0)
		for i := 1; i <= 2*n; i++ {
			if !done[i] {
				rem = append(rem, i)
			}
		}
		sort.Ints(rem)
		m := len(rem)
		for i := 0; i < m/2; i++ {
			chords = append(chords, Chord{rem[i], rem[(m/2)+i]})
		}
		ans := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				chord1 := chords[i]
				chord2 := chords[j]

				if (chord1.p1 > chord2.p1 && chord1.p1 < chord2.p2 && chord1.p2 > chord2.p2) ||
					(chord1.p2 > chord2.p1 && chord1.p2 < chord2.p2 && chord1.p1 < chord2.p1) {
					ans++
				}
			}
		}
		//write(f, chords, "\n")
		write(f, ans, "\n")

	}
}
