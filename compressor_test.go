package archiver

import (
	"os"
	"testing"
)

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
	compareFiles(src, "packed.tar.gz")

	err = Zip(src, dst2)
	if err != nil {
		t.Fatalf("Error compressing \"%s\" to \"%s\": %s", src, dst2, err)
	}
	compareFiles(src, dst2+".gz")
}
