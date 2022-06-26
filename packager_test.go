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

func TestPackFolder1(t *testing.T) {
	folder := "test_assets/testfolder"
	output := "unpackfolder1/packed.tar"

	data, err := PackFolder(folder)
	if err != nil {
		t.Fatalf("Error packaging folder \"%s\" in \"%s\": %s", folder, output, err)
	}

	err = WriteFile(output, data, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", output, err)
	}
	defer os.RemoveAll("unpackfolder1/")

	if !compareFiles(output, "test_assets/testmodels/packed.tar") {
		t.Fatalf("Not valid tar file format")
	}
}

func TestPackFolder2(t *testing.T) {
	folder := "test_assets/testfolder2"
	output := "unpackfolder2/packed.tar"

	data, err := PackFolder(folder)
	if err != nil {
		t.Fatalf("Error packaging folder \"%s\" in \"%s\": %s", folder, output, err)
	}

	err = WriteFile(output, data, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", output, err)
	}
	defer os.RemoveAll("unpackfolder2/")

	if !compareFiles(output, "test_assets/testmodels/packed2.tar") {
		t.Fatalf("Not valid tar file format")
	}
}

/*func TestUnpack(t *testing.T) {
	dst := "tmp"
	src := "testmodels/packed2.tar" //"testmodels/packed.tar"
	//os.Mkdir(dst, 0777)
	input, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file")
	}
	err = Unpack(input, dst)
	if err != nil {
		t.Fatalf("Error unpacking: %s", err.Error())
	}
	/*if !compareFiles("testfolder/samples2/samples21/file6.txt", "tmp/testfolder/samples2/samples21/file6.txt") {
		t.Fatalf("Not valid unpacked files")
	}
	os.RemoveAll("tmp/")
}*/

func TestUnpack1(t *testing.T) {
	dst := "packed1"
	src := "test_assets/testmodels/packed.tar"
	input, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file")
	}
	err = Unpack(input, dst)
	if err != nil {
		t.Fatalf("Error unpacking: %s", err.Error())
	}
	if !compareFiles("test_assets/testfolder/samples2/samples21/file6.txt", "packed1/testfolder/samples2/samples21/file6.txt") {
		t.Fatalf("Not valid unpacked files")
	}
	os.RemoveAll("packed1/")
}

func TestUnpack2(t *testing.T) {
	dst := "packed2"
	src := "test_assets/testmodels/packed2.tar"
	input, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file")
	}
	err = Unpack(input, dst)
	if err != nil {
		t.Fatalf("Error unpacking: %s", err.Error())
	}
	if !compareFiles("test_assets/testfolder2/onlyfile.txt", "packed2/testfolder2/onlyfile.txt") {
		t.Fatalf("Not valid unpacked files")
	}
	os.RemoveAll("packed2/")
}
