package randpics

import (
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func getPics(sourcePath string, count int) []string {
	var allFiles []string
	filepath.Walk(sourcePath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".jpg", f.Name())
			if err == nil && r {
				allFiles = append(allFiles, path)
			}
		}
		return nil
	})

	if len(allFiles) == 0 {
		return allFiles
	}

	rand.Seed(time.Now().UTC().UnixNano())

	var randFiles []string
	for i := 0; i < count; i++ {
		randFiles = append(randFiles, allFiles[rand.Intn(len(allFiles))])
	}

	return randFiles
}
