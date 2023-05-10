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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	//first_answer := 0
	for t := 1; t <= T; t++ {
		var n int
		fmt.Fscan(reader, &n)
		p := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
		}
		adj := make([][]int, n+1)
		for i := 0; i < n; i++ {
			if p[i] != i+1 {
				adj[i+1] = append(adj[i+1], p[i])
			}
		}
		nodes := make([]int, 0)
		for i := 1; i <= n; i++ {
			if len(adj[i]) > 0 {
				nodes = append(nodes, i)
			}
		}
		visited := map[int]bool{}
		ans := 0
		present := false
		for i := 0; i < len(nodes); i++ {
			if visited[nodes[i]] {
				continue
			}
			group := make([]int, 1)
			visited[nodes[i]] = true
			stack := make([]int, 1)
			group[0] = nodes[i]
			stack[0] = nodes[i]
			for len(stack) > 0 {
				ll := len(stack)
				last := stack[ll-1]
				stack = stack[:ll-1]

				for _, v := range adj[last] {
					if visited[v] {
						continue
					}
					stack = append(stack, v)
					group = append(group, v)
					visited[v] = true
				}
			}
			//write(f, group, "\n")
			if present {
				ans += len(group) - 1
			} else {
				sort.Ints(group)
				for j := 1; j < len(group); j++ {
					if group[j] == group[j-1]+1 {
						present = true
						break
					}
				}
				if present {
					ans += len(group) - 2
				} else {
					ans += len(group) - 1
				}
			}
		}
		if !present {
			ans++
		}
		write(f, ans, "\n")
	}
}
