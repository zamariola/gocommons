package file

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

const (
	TESTDATA_PATH = "../testdata"
)

var update = flag.Bool("update", false, "update .golden files")

func TestShouldCopyFileContents(t *testing.T) {
	destination := filepath.Join(TESTDATA_PATH, t.Name()+".tmp")

	os.Remove(destination)

	golden := filepath.Join(TESTDATA_PATH, "sample_source_file.golden")

	if *update || !Exists(golden) {
		fmt.Println("Writing goldenfile " + golden)
		createTestdataIfNoPresent()
		err := ioutil.WriteFile(golden, []byte(fmt.Sprint(rand.Intn(100000))), 0644)

		if err != nil {
			t.Fatal(err)
		}
	}

	Copy(golden, destination)

	expectedFile, _ := ioutil.ReadFile(golden)
	actualFile, _ := ioutil.ReadFile(destination)

	if string(expectedFile) != string(actualFile) {
		t.Errorf("Error while copying file, expected: %s, got: %s", string(expectedFile), string(actualFile))
	}
}

func TestShouldReturnFalseForNonExistingFileDir(t *testing.T) {
	createTestdataIfNoPresent()
	destination := filepath.Join(TESTDATA_PATH, t.Name()+".tmp")
	os.Remove(destination)

	result := Exists(destination)

	if result {
		t.Error("Returned true for non existing file")
	}
}
func TestShouldReturnTrueForExistingFile(t *testing.T) {
	createTestdataIfNoPresent()
	destination := filepath.Join(TESTDATA_PATH, t.Name()+".tmp")
	ioutil.WriteFile(destination, []byte(fmt.Sprint(rand.Intn(100000))), 0644)

	result := Exists(destination)

	if !result {
		t.Error("Returned false for existing file")
	}
}

func TestShouldReturnFalseForNonExistingDir(t *testing.T) {

	result := Exists("/tmp/some/inexisting/path/and/file.tmp")
	if result {
		t.Error("Returned true for non existing dir")
	}
}

func createTestdataIfNoPresent() {
	err := os.MkdirAll(TESTDATA_PATH, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
