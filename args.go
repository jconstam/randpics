package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

var flagSourcePath *string
var flagDestPath *string
var flagCount *int
var excludeFilterList *string

var argsSetup = false

func setupArgs() {
	if !argsSetup {
		flagSourcePath = flag.String("src", "", "The source directory where the pictures are located.")
		flagDestPath = flag.String("dst", "", "The destination directory where the pictures well be transferred.")
		flagCount = flag.Int("count", 0, "The number of pictures to copy.")
		excludeFilterList = flag.String("exclude", "", "A comma-seperated list of words to exclude if they are in the path of a picture.")

		argsSetup = true
	}
}

func parseCommaList(listString string) []string {
	if listString == "" {
		return []string{}
	} else if strings.Contains(listString, ",") {
		return strings.Split(listString, ",")
	} else {
		return []string{listString}
	}
}

func parseArgs() (bool, string, string, int, []string) {
	flag.Parse()

	source, _ := filepath.Abs(*flagSourcePath)
	dest, _ := filepath.Abs(*flagDestPath)
	count := *flagCount
	excludeList := parseCommaList(*excludeFilterList)

	return validateArgs(source, dest, count), *flagSourcePath, *flagDestPath, *flagCount, excludeList
}

func validateArgs(sourcePath string, destPath string, count int) bool {
	return pathExists(sourcePath) && pathExists(destPath) && count > 0
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
