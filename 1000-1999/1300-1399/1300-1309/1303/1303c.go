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
	var T int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &T)
	for t := 1; t <= T; t++ {
		var s string
		fmt.Fscan(reader, &s)
		adj := make([][]byte, 26)
		ok := true
		for i := 1; i < len(s); i++ {
			ch1, ch2 := s[i], s[i-1]
			if len(adj[ch1-'a']) == 0 {
				adj[ch1-'a'] = append(adj[ch1-'a'], ch2)
			}else if len(adj[ch1-'a']) == 1  {
				if ch2 != adj[ch1-'a'][0] {
					adj[ch1-'a'] = append(adj[ch1-'a'], ch2)
				}
			}else if len(adj[ch1-'a']) == 2 {
				if ch2 != adj[ch1-'a'][0] &&  ch2 != adj[ch1-'a'][1] { 
					ok = false
				}
			}

			if len(adj[ch2-'a']) == 0 {
				adj[ch2-'a'] = append(adj[ch2-'a'], ch1)
			}else if len(adj[ch2-'a']) == 1  {
				if ch1 != adj[ch2-'a'][0] {
					adj[ch2-'a'] = append(adj[ch2-'a'], ch1)
				}
			}else if len(adj[ch2-'a']) == 2 {
				if ch1 != adj[ch2-'a'][0] &&  ch1 != adj[ch2-'a'][1] { 
					ok = false
				}
			}
		}
		if ok {
			ans := make([]byte, 0)
			included := make([]bool, 26)
			for i := byte(0); i < 26; i++ {
				if included[i] {
					continue
				}
				if len(adj[i]) == 0 {
					ans = append(ans, 'a'+i)
					included[i] = true
				}else if len(adj[i]) == 1 {
					prev := 'a' + i
					ans = append(ans, prev)
					included[prev-'a'] = true
					curr := adj[prev-'a'][0]
					ans = append(ans, curr)
					included[curr-'a'] = true

					for true {
						if len(adj[curr-'a']) == 1 {
							break
						}else if len(adj[curr-'a']) == 2 {
							if adj[curr-'a'][0] == prev  {
								prev = curr
								curr =  adj[curr-'a'][1]
								ans = append(ans, curr)
								included[curr-'a'] = true
							}else{
								prev = curr
								curr =  adj[curr-'a'][0]
								ans = append(ans, curr)
								included[curr-'a'] = true
							}
						}
					}
					
				}
			}
			if len(ans) == 26 {
				write(f, "YES\n")
				write(f, string(ans),"\n")
			}else{
				ok = false
			}
			
		}
		if !ok{
			write(f, "NO\n")
		}
	}
}
