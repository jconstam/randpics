package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

var setup = false

func parseArgsTest(t *testing.T, argsGoodExpected bool, sourceExpected string, destExpected string, countExpected int, excludeListExpected []string) {
	if !setup {
		setupArgs()
		setup = true
	}

	argsGood, source, dest, count, excludeList := parseArgs()

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
	if len(excludeList) != len(excludeListExpected) {
		t.Errorf("Expected exclude list size to be \"%v\", got \"%v\"", len(excludeListExpected), len(excludeList))
	}

	for index, value := range excludeList {
		if value != excludeListExpected[index] {
			t.Errorf("Expected exclude list at index \"%v\" to contain \"%v\", got \"%v\".", index, value, excludeListExpected[index])
		}
	}
}

func buildFakeArgs(sourcePath string, destPath string, count int, excludeList []string) {
	os.Args = []string{"", "-src=" + sourcePath, "-dst=" + destPath, "-count=" + strconv.Itoa(count), "-exclude=" + strings.Join(excludeList, ",")}
}

func TestParse_BadSource(t *testing.T) {
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs("aa", dest, 1, []string{})

	parseArgsTest(t, false, "aa", dest, 1, []string{})

	os.Remove(dest)
}

func TestParse_BadDest(t *testing.T) {
	source, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, "aa", 1, []string{})

	parseArgsTest(t, false, source, "aa", 1, []string{})

	os.Remove(source)
}
func TestParse_BadCount(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 0, []string{})

	parseArgsTest(t, false, source, dest, 0, []string{})

	os.Remove(source)
	os.Remove(dest)
}
func TestParse_Good(t *testing.T) {
	source, _ := ioutil.TempDir("", "")
	dest, _ := ioutil.TempDir("", "")

	buildFakeArgs(source, dest, 1, []string{})

	parseArgsTest(t, true, source, dest, 1, []string{})

	os.Remove(source)
	os.Remove(dest)
}
