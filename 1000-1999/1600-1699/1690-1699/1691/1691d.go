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

type Pair struct {
	val int64
	i   int
}
type Range struct {
	sum  int64
	i, j int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}
		// if T > 1200 {
		// 	if t == 1732 {
		// 		fmt.Println(a)

		// 	} else {
		// 		continue
		// 	}
		// }
		b := make([]int64, 0)
		found := false
		for i := 0; i < n; i++ {
			if i == 0 {
				b = append(b, a[i])
			} else {
				last := len(b) - 1
				if b[last] > 0 {
					if a[i] > 0 {
						found = true
						break
					} else if a[i] == 0 {

					} else {
						b = append(b, a[i])
					}
				} else {
					if a[i] > 0 {
						b = append(b, a[i])
					} else {
						b[last] += a[i]
					}
				}
			}
		}
		//write(f, b, found, "\n")
		if !found {
			m := len(b)
			pairs := make([]Pair, m)
			for i := 0; i < m; i++ {
				pairs[i] = Pair{val: b[i], i: i}
			}
			sort.Slice(pairs, func(i, j int) bool { return pairs[i].val < pairs[j].val })
			//write(f, pairs, "\n")
			ranges := make([]Range, m)
			for i := 0; i < m; i++ {
				pair := pairs[i]
				if pair.val <= 0 {
					ranges[pair.i] = Range{i: pair.i, j: pair.i, sum: pair.val}
				} else {
					if pair.i > 0 && (ranges[pair.i-1].sum > 0 || ranges[pair.i-1].i != 0 && ranges[pair.i-1].sum+pair.val > 0) || pair.i < m-1 && (ranges[pair.i+1].sum > 0 || ranges[pair.i+1].j != m-1 && ranges[pair.i+1].sum+pair.val > 0) {
						found = true
						break
					} else {
						min_index := pair.i
						max_index := pair.i
						sum := pair.val
						if pair.i > 0 {
							min_index = min(min_index, ranges[pair.i-1].i)
							sum += ranges[pair.i-1].sum
						}
						if pair.i < m-1 {
							max_index = max(max_index, ranges[pair.i+1].j)
							sum += ranges[pair.i+1].sum
						}
						ranges[max_index] = Range{i: min_index, j: max_index, sum: sum}
						ranges[min_index] = ranges[max_index]
					}
				}
			}
			//	write(f, ranges, "\n")

		}
		if found {
			write(f, "NO\n")
		} else {
			write(f, "YES\n")
		}
	}

}
