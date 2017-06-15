package main

import (
	"fmt"

	"github.com/jconstam/randpics/args"
	"github.com/jconstam/randpics/picDest"
	"github.com/jconstam/randpics/picSource"
)

func main() {
	argsGood, source, dest, count := args.ParseArgs()
	if argsGood != true {
		fmt.Printf("ERROR PARSING ARGS: %v %v %v\n", source, dest, count)
		return
	}

	fmt.Printf("Searching in %v for pictures...\n", source)
	sourcePics := picSource.GetPics(source, count)

	fmt.Printf("Moving %v pictures to %v...\n", count, dest)
	picDest.MoveToDest(dest, sourcePics)

	fmt.Printf("Done!\n")
}
