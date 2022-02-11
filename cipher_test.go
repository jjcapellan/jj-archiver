package archiver

import "testing"

func TestEncrypt(t *testing.T) {
	src := "testmodels/packed.tar.gz"
	dst := ""
	password := "axdcf"

	err := Encrypt(src, dst, password)
	if err != nil {
		t.Fatalf("Encryption error: %s", err.Error())
	}
}
