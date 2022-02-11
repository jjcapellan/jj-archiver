package archiver

import (
	"os"
	"testing"
)

func TestDecrypt(t *testing.T) {
	src := "testmodels/packed2.tar.gz.crp"
	dst := ""
	password := "axdcf"
	defer os.Remove("packed2.tar.gz")

	err := Decrypt(src, dst, password)
	if err != nil {
		t.Fatalf("Decryption error: %s", err.Error())
	}
	if !compareFiles("testmodels/packed.tar.gz", "packed2.tar.gz") {
		t.Fatalf("Not valid decrypted file")
	}
}

func TestEncrypt(t *testing.T) {
	src := "testmodels/packed.tar.gz"
	dst := ""
	password := "axdcf"

	err := Encrypt(src, dst, password)
	if err != nil {
		t.Fatalf("Encryption error: %s", err.Error())
	}
}
