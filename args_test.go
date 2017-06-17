package randpics

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

var setup = false

func parseArgsTest(t *testing.T, argsGoodExpected bool, sourceExpected string, destExpected string, countExpected int) {
	if !setup {
		setupArgs()
		setup = true
	}

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

func TestParse_BadSource(t *testing.T) {
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs("aa", dest, 1)

	parseArgsTest(t, false, "aa", dest, 1)
}

func TestParse_BadDest(t *testing.T) {
	source, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, "aa", 1)

	parseArgsTest(t, false, source, "aa", 1)
}
func TestParse_BadCount(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 0)

	parseArgsTest(t, false, source, dest, 0)
}
func TestParse_Good(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 1)

	parseArgsTest(t, true, source, dest, 1)
}
