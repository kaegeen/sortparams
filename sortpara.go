package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	n := len(args)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}
