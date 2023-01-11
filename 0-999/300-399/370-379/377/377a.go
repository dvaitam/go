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

var last []int

func dfs(s [][]byte, visited [][]bool, n int, m int, i int, j int) {
	if s[i][j] != '.' {
		return
	}
	visited[i][j] = true
	last = append(last, i*m+j)
	if i+1 < n && !visited[i+1][j] {
		dfs(s, visited, n, m, i+1, j)
	}
	if j+1 < m && !visited[i][j+1] {
		dfs(s, visited, n, m, i, j+1)
	}
	if i > 0 && !visited[i-1][j] {
		dfs(s, visited, n, m, i-1, j)
	}
	if j > 0 && !visited[i][j-1] {
		dfs(s, visited, n, m, i, j-1)
	}
}

//	func broken(s [][]byte, n int, m int) bool {
//		visited := make([][]bool, n)
//		for i := 0; i < n; i++ {
//			visited[i] = make([]bool, m)
//		}
//		count := 0
//		for i := 0; i < n; i++ {
//			for j := 0; j < m; j++ {
//				if s[i][j] == '.' && !visited[i][j] {
//					if count > 0 {
//						return true
//					}
//					dfs(s, visited, n, m, i, j)
//					count++
//				}
//			}
//		}
//		return false
//	}
func get_hash_count(s [][]byte, n int, m int, i int, j int) int {
	count := 0
	rows := []int{i - 1, i, i + 1}
	cols := []int{j - 1, j, j + 1}
	for l := 0; l < 3; l++ {
		for ll := 0; ll < 3; ll++ {
			row, col := rows[l], cols[ll]
			if row >= 0 && col >= 0 && row < n && col < m {
				if s[row][col] == '#' {
					count++
				}
			}
		}
	}
	return count
}
func get_connected_count(s [][]byte, n int, m int, i int, j int) int {
	count := 0
	if i+1 < n && s[i+1][j] == '.' {
		count++
	}
	if j+1 < m && s[i][j+1] == '.' {
		count++
	}
	if i > 0 && s[i-1][j] == '.' {
		count++
	}
	if j > 0 && s[i][j-1] == '.' {
		count++
	}

	return count
}

func main() {
	var n, m, k int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n, &m, &k)
	s := make([]string, n)
	sm := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		sm[i] = []byte(s[i])
	}
	last = make([]int, 0)
	count := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n && count == 0; i++ {
		for j := 0; j < m && count == 0; j++ {
			if sm[i][j] != '.' {
				continue
			}
			count++
			dfs(sm, visited, n, m, i, j)

		}
	}
	//write(f, visited, last, "\n")
	start := len(last) - 1
	for k > 0 {
		i, j := last[start]/m, last[start]%m
		sm[i][j] = 'X'
		k--
		start--
	}

	for i := 0; i < n; i++ {
		write(f, string(sm[i]), "\n")
	}

}
