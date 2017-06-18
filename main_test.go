package randpics

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func createTempPicFile(sourceDir string) string {
	tempFile, _ := ioutil.TempFile(sourceDir, "")
	tempFile.Close()
	os.Rename(tempFile.Name(), tempFile.Name()+".jpg")

	return filepath.Base(tempFile.Name() + ".jpg")
}

func TestMain_CopyAllNoDelete(t *testing.T) {
	dest, _ := ioutil.TempDir("", "")
	source, _ := ioutil.TempDir("", "")

	file1 := createTempPicFile(source)

	buildFakeArgs(source, dest, 1)
	main()

	if _, err := os.Stat(filepath.Join(dest, file1)); os.IsNotExist(err) {
		t.Errorf("Couldn't find copied file %v in destination path %v", file1, dest)
	}

	os.Remove(source)
	os.Remove(dest)
}
