package main

import (
	"fmt"

	"github.com/jconstam/randpics/args"
)

func main() {
	argsGood, source, dest, count := args.ParseArgs()
	if argsGood != true {
		fmt.Printf("ERROR PARSING ARGS: %v %v %v\n", source, dest, count)
		return
	}
}
