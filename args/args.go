package args

import (
	"flag"
	"os"
	"path/filepath"
)

func ParseArgs() (bool, string, string, int) {
	sourcePath := flag.String("src", "", "The source directory where the pictures are located.")
	destPath := flag.String("dst", "", "The destination directory where the pictures well be transferred.")
	count := flag.Int("count", 0, "The number of pictures to copy.")

	flag.Parse()

	source, _ := filepath.Abs(*sourcePath)
	dest, _ := filepath.Abs(*destPath)

	return validateArgs(source, dest, *count), *sourcePath, *destPath, *count
}

func validateArgs(sourcePath string, destPath string, count int) bool {
	return pathExists(sourcePath) && pathExists(destPath) && count > 0
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
