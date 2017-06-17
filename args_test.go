package randpics

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func testParse(t *testing.T, argsGoodExpected bool, sourceExpected string, destExpected string, countExpected int) {
	argsGood, source, dest, count := parseArgs()

	if argsGood != argsGoodExpected {
		t.Errorf("Expected parse status to be \"%v\", got \"%v\"", argsGoodExpected, argsGood)
	}
	if source != sourceExpected {
		t.Errorf("Expected source to be \"%v\", got \"%v\"", sourceExpected, source)
	}
	if dest != destExpected {
		t.Errorf("Expected dest to be \"%v\", got \"%v\"", destExpected, dest)
	}
	if count != countExpected {
		t.Errorf("Expected count to be \"%v\", got \"%v\"", countExpected, count)
	}
}

func buildFakeArgs(sourcePath string, destPath string, count int) {
	os.Args = []string{"", "-src=" + sourcePath, "-dst=" + destPath, "-count=" + strconv.Itoa(count)}
}

/*
func TestParse_BadSource(t *testing.T) {
	//source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs("aa", dest, 1)

	testParse(t, false, "aa", dest, 1)
}
func TestParse_BadDest(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	//dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, "aa", 1)

	testParse(t, false, source, "aa", 1)
}
func TestParse_BadCount(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 0)

	testParse(t, false, source, dest, 0)
}
*/
func TestParse_Good(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 1)

	testParse(t, true, source, dest, 1)
}
