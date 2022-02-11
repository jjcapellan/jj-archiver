package archiver

import (
	"bytes"
	"os"
	"testing"
)

// compareFiles returns true if file1 == file2
func compareFiles(file1 string, file2 string) bool {

	b1, _ := readFile(file1)
	b2, _ := readFile(file2)

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
	output1 := "packed"
	output2 := "unpackfolder1/packed"
	output3 := "unpackfolder2/subfolder/packed"
	defer os.Remove("packed.tar")
	defer os.RemoveAll("unpackfolder1/")
	defer os.RemoveAll("unpackfolder2/")

	err := PackFolder(folder, output1)
	if err != nil {
		t.Fatalf("Error packaging folder \"testfolder\" in \"%s\": %s", output1+"tar", err)
	}

	if !compareFiles(output1+".tar", "testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}

	err = PackFolder(folder, output2)
	if err != nil {
		t.Fatalf("Error packaging folder \"testfolder\" in \"%s\": %s", output1+"tar", err)
	}

	if !compareFiles(output2+".tar", "testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}

	err = PackFolder(folder, output3)
	if err != nil {
		t.Fatalf("Error packaging folder \"testfolder\" in \"%s\": %s", output3+"tar", err)
	}

	if !compareFiles(output3+".tar", "testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}
}

func TestUnpack(t *testing.T) {
	dst := "tmp"
	src := "testmodels/packed.tar"
	os.Mkdir(dst, 0777)
	err := Unpack(src, dst)
	if err != nil {
		t.Fatalf("Error unpacking")
	}
	if !compareFiles("testfolder/samples2/samples21/file6.txt", "tmp/testfolder/samples2/samples21/file6.txt") {
		t.Fatalf("Not valid unpacked files")
	}
	os.RemoveAll("tmp/")
}
