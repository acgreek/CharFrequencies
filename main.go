// You can edit this code!
// Click here and start typing.
package main

import (
		"os"
		"fmt"
		"bufio"
		"sort"
		)

func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
			err error = nil
			line, ln []byte
			)
		for isPrefix && err == nil {
			line, isPrefix, err = r.ReadLine()
				ln = append(ln, line...)
		}
	return string(ln),err
}

func main() {
	

	var x map[byte ] int = make(map[byte]int)
	r := bufio.NewReader(os.Stdin)
	s, e := Readln(r)
	for e == nil {
		for i := range s {
			if _, ok := x[s[i] ]; ok {
				x[s[i] ] = x[s[i]] + 1
			} else {
				x[s[i] ] = 1 
			}
		}
		s,e = Readln(r)
	}

	n := map[int][]byte{}
	var a []int
		for k, v := range x {
			n[v] = append(n[v], k)
		}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
		for _, k := range a {
			for _, s := range n[k] {
				fmt.Printf("%c, %d\n", s, k)
			}
		}
}
