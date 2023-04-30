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
func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

type Edge struct {
	u, v int
}

func main() {
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		adj := map[int][]int{}
		degree := make([]int, n+1)
		edges := make([]Edge, m)
		for i := 0; i < m; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			degree[u]++
			degree[v]++
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
			edges[i] = Edge{u: u, v: v}
		}
		// if T > 140 {
		// 	if t == 60 {
		// 		fmt.Println(adj)
		// 	} else {
		// 		continue
		// 	}
		// }
		ok := false
		visited := make([]bool, n+1)
		parent := make([]int, n+1)

		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}
			//fmt.Println("exploring", i)
			stack := make([]int, 1)
			stack[0] = i

			visited[i] = true

			for len(stack) > 0 {
				ll := len(stack)
				last := stack[ll-1]
				stack = stack[:ll-1]
				//fmt.Println("goinging down", last, stack)
				for _, v := range adj[last] {
					if visited[v] {
						continue
					}
					//fmt.Println("visitted", v, stack)
					visited[v] = true
					stack = append(stack, v)
					parent[v] = last
				}
			}
		}
		//fmt.Println(parent)
		//write(f, parent, "\n")
		for i := 0; i < m; i++ {
			parents := map[int]bool{}
			edge := edges[i]
			u := edge.u
			parents[u] = true
			for parent[u] != 0 {
				parents[parent[u]] = true
				u = parent[u]
			}
			//fmt.Println(parents)
			v := edge.v
			common := -1
			for parent[v] != 0 {
				if parents[parent[v]] && v != edge.u && !parents[v] {
					common = parent[v]
					break
				}
				v = parent[v]
			}
			// if edge.u == 3 && edge.v == 2 || edge.u == 2 && edge.v == 3 {
			// 	fmt.Println("cl", common, parents)
			// }
			if common != -1 {
				//fmt.Println("found", common, edge)
				if common == edge.u || common == edge.v {
					continue
				}
				nodes := make([]int, 1)
				nodes[0] = edge.u
				u := edge.u
				for parent[u] != common {
					nodes = append(nodes, parent[u])
					u = parent[u]
				}
				reverse(nodes)
				nodes = append(nodes, edge.v)
				v := edge.v
				for parent[v] != common {
					nodes = append(nodes, parent[v])
					v = parent[v]
				}
				nodes = append(nodes, common)
				ans := make([]Edge, 0)
				nodes_map := map[int]bool{}
				for j := 0; j < len(nodes); j++ {
					nodes_map[nodes[j]] = true
				}
				for j := 0; j < len(nodes); j++ {
					if degree[nodes[j]] > 3 {
						for _, v := range adj[nodes[j]] {
							if !nodes_map[v] {
								ans = append(ans, Edge{u: nodes[j], v: v})
							}
							if len(ans) == 2 {
								break
							}
						}
						break
					}
				}
				if len(ans) > 1 {
					for j := 0; j < len(nodes); j++ {
						ans = append(ans, Edge{nodes[j], nodes[(j+1)%len(nodes)]})
					}
					ok = true
					write(f, "YES\n")
					write(f, len(ans), "\n")
					for j := 0; j < len(ans); j++ {
						write(f, ans[j].u, ans[j].v, "\n")
					}
					break
				}
			}
		}
		if !ok {
			write(f, "NO\n")
		}
	}

}
