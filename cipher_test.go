package archiver

import (
	"os"
	"testing"
)

func TestDecrypt(t *testing.T) {
	src := "testmodels/packed.tar.gz.crp"
	dst := "testdecrypt"
	password := "axdcf"
	defer os.RemoveAll("testdecrypt/")

	err := Decrypt(src, dst, password)
	if err != nil {
		t.Fatalf("Decryption error: %s", err.Error())
	}
	if !compareFiles("testmodels/packed.tar.gz", "testdecrypt/packed.tar.gz") {
		t.Fatalf("Not valid decrypted file")
	}
}

func TestEncrypt(t *testing.T) {
	src := "testmodels/packed.tar.gz"
	dst := "testencrypt"
	password := "axdcf"
	defer os.RemoveAll("testencrypt/")

	err := Encrypt(src, dst, password)
	if err != nil {
		t.Fatalf("Encryption error: %s", err.Error())
	}

	err = Decrypt(dst+"/packed.tar.gz.crp", dst, password)
	if err != nil {
		t.Fatalf("Decryption error: %s", err.Error())
	}

	if !compareFiles("testmodels/packed.tar.gz", dst+"/packed.tar.gz") {
		t.Fatalf("Not valid encrypted file")
	}
}
