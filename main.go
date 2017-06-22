// You can edit this code!
// Click here and start typing.
package main

import (
		"os"
		"fmt"
		)

func main() {
	data := make([]byte, 100)
	count, _:= os.Stdin.Read(data)
	var x map[byte ] int = make(map[byte]int)
	for i := 0; i < count; i++  {
			if _, ok := x[data[i] ]; ok {
		  	x[data[i] ] = x[data[i]] + 1
			} else {
		  	x[data[i] ] = 1 
			}
	}
	for k, _ := range x {
		fmt.Printf("%c=%d\n", k, x[k])
	}
}
