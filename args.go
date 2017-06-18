package randpics

import (
	"flag"
	"os"
	"path/filepath"
)

var flagSourcePath *string
var flagDestPath *string
var flagCount *int

var argsSetup = false

func setupArgs() {
	if !argsSetup {
		flagSourcePath = flag.String("src", "", "The source directory where the pictures are located.")
		flagDestPath = flag.String("dst", "", "The destination directory where the pictures well be transferred.")
		flagCount = flag.Int("count", 0, "The number of pictures to copy.")

		argsSetup = true
	}
}

func parseArgs() (bool, string, string, int) {
	flag.Parse()

	source, _ := filepath.Abs(*flagSourcePath)
	dest, _ := filepath.Abs(*flagDestPath)
	count := *flagCount

	return validateArgs(source, dest, count), *flagSourcePath, *flagDestPath, *flagCount
}

func validateArgs(sourcePath string, destPath string, count int) bool {
	return pathExists(sourcePath) && pathExists(destPath) && count > 0
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
