package randpics

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTempPicFile(sourceDir string) {
	tempFile, _ := ioutil.TempFile(sourceDir, "")
	tempFile.Close()
	os.Rename(tempFile.Name(), tempFile.Name()+".jpg")
}

func TestMain_CopyAllNoDelete(t *testing.T) {
	dest, _ := ioutil.TempDir("", "")
	source, _ := ioutil.TempDir("", "")

	createTempPicFile(source)

	buildFakeArgs(source, dest, 1)

	main()

	os.Remove(source)
	os.Remove(dest)
}
