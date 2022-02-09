package archiver

import (
	"os"
	"testing"
)

// compareFiles returns true if file1 == file2
func compareFiles(file1 string, file2 string) bool {
	buffer1, _ := readFile(file1)
	buffer2, _ := readFile(file2)

	if len(buffer1) != len(buffer2) {
		return false
	}

	for i, byte := range buffer1 {
		if byte != buffer2[i] {
			return false
		}
	}

	return true
}

func TestPackFolder(t *testing.T) {
	folder := "testfolder"
	output := "packed"
	defer os.Remove(output + ".tar")

	err := PackFolder(folder, output)
	if err != nil {
		t.Fatalf("Error packaging folder: %s", err)
	}

	if !compareFiles("packed.tar", "testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}
}
