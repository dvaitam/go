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
func get_prob(w int, b int) float64 {
	if w == 0 || b < 0 {
		return 0
	}
	if b == 0 {
		return 1
	}

	ans := float64(w) / float64(w+b)
	if b == 1 {
		return ans
	}
	ans += (float64(b*(b-1)*(b-2))*get_prob(w, b-3) + float64(b*(b-1)*w)*get_prob(w-1, b-2)) / float64((w+b)*(w+b-1)*(w+b-2))
	return ans
}
func main() {
	var w, b int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &w, &b)
	dp := make([][]float64, b+1)
	for i := 0; i <= b; i++ {
		dp[i] = make([]float64, w+1)
	}
	for bi := 0; bi < min(3, b+1); bi++ {
		for wi := 0; wi <= w; wi++ {
			dp[bi][wi] = get_prob(wi, bi)
		}
	}
	for wi := 0; wi < min(2, w+1); wi++ {
		for bi := 0; bi <= b; bi++ {
			dp[bi][wi] = get_prob(wi, bi)
		}
	}
	for bi := 3; bi <= b; bi++ {
		for wi := 2; wi <= w; wi++ {
			dp[bi][wi] = float64(wi) / float64(wi+bi)
			dp[bi][wi] += (float64(bi*(bi-1)*(bi-2))*dp[bi-3][wi] + float64(bi*(bi-1)*wi)*dp[bi-2][wi-1]) / float64((wi+bi)*(wi+bi-1)*(wi+bi-2))
		}
	}
	write(f, dp[b][w], "\n")
}
