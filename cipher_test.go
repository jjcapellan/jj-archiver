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
	dst := "testencrypt/packed.tar.gz.crp"
	password := "axdcf"
	defer os.RemoveAll("testencrypt/")

	srcData, err := ReadFile(src)
	if err != nil {
		t.Fatalf("Error reading file %s: %s", src, err)
	}

	outData, err := Encrypt(srcData, password)
	if err != nil {
		t.Fatalf("Encryption error: %s", err.Error())
	}

	err = WriteFile(dst, outData, 0666)
	if err != nil {
		t.Fatalf("Error writing file %s : %s", dst, err)
	}

	err = Decrypt(dst, "testencrypt", password)
	if err != nil {
		t.Fatalf("Decryption error: %s", err.Error())
	}

	if !compareFiles("testmodels/packed.tar.gz", "testencrypt/packed.tar.gz") {
		t.Fatalf("Not valid encrypted file")
	}
}
