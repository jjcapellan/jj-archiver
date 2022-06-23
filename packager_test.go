package archiver

import (
	"bytes"
	"os"
	"testing"
)

// compareFiles returns true if file1 == file2
func compareFiles(file1 string, file2 string) bool {

	b1, _ := ReadFile(file1)
	b2, _ := ReadFile(file2)

	if len(b1) != len(b2) {
		return false
	}

	if bytes.Compare(b1, b2) != 0 {
		return false
	}

	return true
}

func TestPackFolder(t *testing.T) {
	folder := "testfolder"
	output1 := "unpackfolder1/packed.tar"
	folder2 := "testfolder/samples2"
	output2 := "unpackfolder2/packed.tar"

	data, err := PackFolder(folder)
	if err != nil {
		t.Fatalf("Error packaging folder \"testfolder\" in \"%s\": %s", output1, err)
	}

	err = WriteFile(output1, data, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", output1, err)
	}
	defer os.RemoveAll("unpackfolder1/")

	if !compareFiles(output1, "testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}

	data, err = PackFolder(folder2)
	if err != nil {
		t.Fatalf("Error packaging folder \"testfolder/samples2\" in \"%s\": %s", output2, err)
	}

	err = WriteFile(output2, data, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", output2, err)
	}
	defer os.RemoveAll("unpackfolder2/")
}

func TestUnpack(t *testing.T) {
	dst := "tmp"
	src := "testmodels/packed.tar"
	os.Mkdir(dst, 0777)
	input, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file")
	}
	err = Unpack(input, dst)
	if err != nil {
		t.Fatalf("Error unpacking")
	}
	if !compareFiles("testfolder/samples2/samples21/file6.txt", "tmp/testfolder/samples2/samples21/file6.txt") {
		t.Fatalf("Not valid unpacked files")
	}
	os.RemoveAll("tmp/")
}
