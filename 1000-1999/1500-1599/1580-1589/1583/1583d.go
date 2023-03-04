// 00
package main

import (
	"bufio"
	"fmt"
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

func print_one_array(j int, n int, val int) {
	other := 1
	if val == 1 {
		other = 2
	}
	fmt.Print("?")
	for i := 1; i <= n; i++ {
		fmt.Print(" ")
		if i == j {
			fmt.Print(other)
		} else {
			fmt.Print(val)
		}
	}
	fmt.Println()
}
func main() {
	var n, k int
	//reader := bufio.NewReader(os.Stdin)
	// f := bufio.NewWriter(os.Stdout)
	// defer f.Flush()
	//fmt.Fscan(reader, &T)
	fmt.Scan(&n)
	greater := map[int]int{}
	lesser := map[int]int{}
	one_index := 0
	n_index := 0
	for i := n; i > 0; i-- {
		print_one_array(i, n, 1)
		fmt.Scan(&k)
		if k == 0 {
			n_index = i
		} else if k != i {
			greater[i] = k
			lesser[k] = i
		}
		print_one_array(i, n, 2)
		fmt.Scan(&k)
		if k == 0 {
			one_index = i
		} else if k != i {
			lesser[i] = k
			greater[k] = i
		}
	}
	ans := make([]int, n+1)
	ans[one_index] = 1
	ans[n_index] = n
	curr := 2
	curr_index := one_index
	for greater[curr_index] > 0 {
		ans[greater[curr_index]] = curr
		curr++
		curr_index = greater[curr_index]
	}
	fmt.Print("!")
	for i := 1; i <= n; i++ {
		fmt.Print(" ", ans[i])
	}
	fmt.Println()
}
