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
	dst := "testzip"
	defer os.RemoveAll("testzip/")

	err := Zip(src, dst)
	if err != nil {
		t.Fatalf("Error compressing \"%s\" to \"%s\": %s", src, dst, err)
	}
	if !compareFiles(src+".gz", dst+"/packed.tar.gz") {
		t.Fatalf("Not valid gzip format")
	}
}
