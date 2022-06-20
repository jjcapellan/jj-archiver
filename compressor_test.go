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
	dst := "testzip/packed.tar.gz"
	defer os.RemoveAll("testzip/")

	srcData, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file %s: %s", src, err)
	}

	dstData, err := Compress(srcData, src)
	if err != nil {
		t.Fatalf("Error compressing %s: %s", src, err)
	}

	err = WriteFile(dst, dstData, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", dst, err)
	}

	if !compareFiles(src+".gz", dst) {
		t.Fatalf("Not valid gzip format")
	}
}

func TestGetDecompressedSize(t *testing.T) {
	fName := "testmodels/packed.tar.gz"
	gotSize, err := GetDecompressedSize(fName)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	f, err := os.Open("testmodels/packed.tar")
	defer f.Close()
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	fInfo, _ := f.Stat()

	expectedSize := fInfo.Size()

	if expectedSize != int64(gotSize) {
		t.Fatalf("Expected size: %d Got size: %d", expectedSize, gotSize)
	}
}
