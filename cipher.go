package archiver

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func Encrypt(src string, dstDir string, password string) error {
	buffer, err := readFile(src)
	if err != nil {
		return err
	}

	gcm, err := getGCM(password)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	result := gcm.Seal(nonce, nonce, buffer, nil)

	dst, err := prepareDst(src, dstDir, ".crp", false)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, result, 0666)
	if err != nil {
		return err
	}

	return nil
}

func Decrypt(src string, dstDir string, password string) error {
	if filepath.Ext(src) != ".crp" {
		return errors.New("Unrecognized file extension")
	}

	buffer, err := readFile(src)
	if err != nil {
		return err
	}

	gcm, err := getGCM(password)
	if err != nil {
		return err
	}

	s := gcm.NonceSize()

	nonce, cipherContent := buffer[:s], buffer[s:]

	content, err := gcm.Open(nil, nonce, cipherContent, nil)
	if err != nil {
		return err
	}

	dst, err := prepareDst(src, dstDir, ".crp", true)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, content, 0777)
	if err != nil {
		return err
	}

	return nil
}

//// HELPERS

func getGCM(password string) (cipher.AEAD, error) {

	// Gets 32 bytes long key from password
	h := sha256.New()
	h.Write([]byte(password))
	key := h.Sum(nil)

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}
