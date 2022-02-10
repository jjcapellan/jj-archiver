package archiver

import (
	"os"
	"testing"
)

func TestUnZip(t *testing.T) {
	src := "testmodels/packed.tar.gz"
	dst1 := ""
	defer os.Remove("packed.tar")

	err := UnZip(src, dst1)
	if err != nil {
		t.Fatalf("Error uncompressing \"%s\" to \"%s\": %s", src, dst1, err)
	}
	if !compareFiles("testmodels/packed.tar", "packed.tar") {
		t.Fatalf("Not valid uncompressed file")
	}
}

func TestZip(t *testing.T) {
	src := "testmodels/packed.tar"
	dst1 := ""
	dst2 := "zipfolder/subfolder/packed"
	defer os.Remove("packed.tar.gz")
	defer os.RemoveAll("zipfolder/")

	err := Zip(src, dst1)
	if err != nil {
		t.Fatalf("Error compressing \"%s\" to \"%s\": %s", src, dst1, err)
	}
	if !compareFiles(src+".gz", "packed.tar.gz") {
		t.Fatalf("Not valid gzip format")
	}

	err = Zip(src, dst2)
	if err != nil {
		t.Fatalf("Error compressing \"%s\" to \"%s\": %s", src, dst2, err)
	}
	if !compareFiles(src+".gz", dst2+".gz") {
		t.Fatalf("Not valid gzip format")
	}
}
