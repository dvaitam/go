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
func get_not(s1 string, s2 string) int {
	mi := min(len(s1), len(s2))
	if s1[:mi] == s2[:mi] && s2 > s1 {
		return -2
	}
	for i := 0; i < mi; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return -1

}
func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	fmt.Fscan(reader, &n)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	possible := true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			possible = false
			break
		}
	}
	if possible {
		write(f, "abcdefghijklmnopqrstuvwxyz\n")
	} else {
		adj := map[byte]map[byte]bool{}
		radj := map[byte]map[byte]bool{}

		inc := map[byte]int{}
		out := map[byte]int{}
		impossible := false
		for i := 1; i < n; i++ {
			j := get_not(s[i], s[i-1])

			if j >= 0 {
				if adj[s[i-1][j]] == nil {
					adj[s[i-1][j]] = map[byte]bool{}
				}
				if radj[s[i][j]] == nil {
					radj[s[i][j]] = map[byte]bool{}
				}
				if !adj[s[i-1][j]][s[i][j]] {
					adj[s[i-1][j]][s[i][j]] = true
					radj[s[i][j]][s[i-1][j]] = true
					inc[s[i][j]]++
					out[s[i-1][j]]++
				}

			} else if j == -2 {
				impossible = true
			}
		}
		// if s[0] == "qqqyoujtewnjomtynhaaiojspmljpt" {
		// 	write(f, inc, out, "\n", adj[byte(101)], "\n")
		// 	write(f, radj[byte(101)], "\n")
		// }
		//write(f, inc, "\n", out, "\n", adj, "\n")
		ans := make([]byte, 0)
		rev := make([]byte, 0)
		l := make([]byte, 0)
		included := map[byte]bool{}
		if !impossible {
			for i := 0; i < 26; i++ {
				c := 'a' + byte(i)
				if inc[c] == 0 || out[c] == 0 {
					l = append(l, c)
					included[c] = true
				}
			}
			//loop := 0

			for len(l) > 0 {
				//write(f, l, "\n")
				curr := l[0]
				//write(f, "dealing with", curr, "\n")
				l = l[1:]
				if inc[curr] == 0 && out[curr] == 0 {
					//write(f, "no in or out ", curr, "\n")
					ans = append(ans, curr)
				} else if inc[curr] == 0 {
					//write(f, "no in ", curr, "\n")

					ans = append(ans, curr)
					remove := make([]byte, 0)
					for k, v := range adj[curr] {
						if v {
							remove = append(remove, k)
						}
					}
					for _, k := range remove {
						adj[curr][k] = false
						radj[k][curr] = false
						out[curr]--
						inc[k]--
						if inc[k] == 0 {
							if !included[k] {
								l = append(l, k)
								included[k] = true
							}
						}
					}
				} else if out[curr] == 0 {
					//write(f, "no out ", curr, "\n")

					rev = append(rev, curr)
					remove := make([]byte, 0)

					for k, v := range radj[curr] {
						if v {
							remove = append(remove, k)

						}
					}
					for _, k := range remove {

						radj[curr][k] = false
						adj[k][curr] = false
						inc[curr]--
						out[k]--
						if out[k] == 0 {
							if !included[k] {
								l = append(l, k)
								included[k] = true
							}
						}
					}
				}
				// loop++
				// if loop > 3 {
				// 	break
				// }
			}

			for i := len(rev) - 1; i >= 0; i-- {
				ans = append(ans, rev[i])
			}
		}

		//write(f, inc, "\n")
		//write(f, out, "\n")
		//write(f, adj, "\n")

		//write(f, ans, len(ans), "\n")
		// if s[0] == "qqqyoujtewnjomtynhaaiojspmljpt" {
		// 	write(f, inc, out, "\n", adj[byte(101)], "\n")
		// 	write(f, radj[byte(101)], "\n")
		// }
		if len(ans) != 26 {
			write(f, "Impossible\n")
		} else {
			write(f, string(ans), "\n")
		}

	}

}
