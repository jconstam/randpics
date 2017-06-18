package randpics

import (
	"fmt"
)

func main() {
	setupArgs()
	argsGood, source, dest, count := parseArgs()
	if argsGood != true {
		fmt.Printf("ERROR PARSING ARGS: %v %v %v\n", source, dest, count)
		return
	}

	fmt.Printf("Searching in %v for pictures...\n", source)
	sourcePics := getPics(source, count)
	if len(sourcePics) == 0 {
		fmt.Printf("No pictures found in %v\n", source)
		return
	}

	fmt.Printf("Moving %v pictures to %v...\n", count, dest)
	moveToDest(dest, sourcePics)

	fmt.Printf("Done!\n")
}
