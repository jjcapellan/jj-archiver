package archiver

import (
	"os"
	"testing"
)

func TestDecompress(t *testing.T) {
	src := "testmodels/packed.tar.gz"
	dst1 := ""
	defer os.Remove("packed.tar")

	err := Decompress(src, dst1)
	if err != nil {
		t.Fatalf("Error decompressing \"%s\" to \"%s\": %s", src, dst1, err)
	}
	if !compareFiles("testmodels/packed.tar", "packed.tar") {
		t.Fatalf("Not valid decompressed file")
	}
}

func TestCompress(t *testing.T) {
	src := "testmodels/packed.tar"
	dst := "testzip"
	defer os.RemoveAll("testzip/")

	err := Compress(src, dst)
	if err != nil {
		t.Fatalf("Error compressing \"%s\" to \"%s\": %s", src, dst, err)
	}
	if !compareFiles(src+".gz", dst+"/packed.tar.gz") {
		t.Fatalf("Not valid gzip format")
	}
}
