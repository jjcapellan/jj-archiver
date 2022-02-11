package archiver

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"os"
	"path/filepath"
)

func Encrypt(src string, dst string, password string) error {
	buffer, err := readFile(src)
	if err != nil {
		return err
	}

	// Gets 32 bytes long key from password
	h := sha256.New()
	h.Write([]byte(password))
	key := h.Sum(nil)

	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	result := gcm.Seal(nonce, nonce, buffer, nil)

	dir, fName := filepath.Split(dst)
	if dst == "" {
		_, fName = filepath.Split(src)
		dst = fName + ".crp"
	} else {
		dst = filepath.Join(dir, fName) + ".crp"
	}
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0777)
		}
	}

	err = os.WriteFile(dst, result, 0666)
	if err != nil {
		return err
	}

	return nil
}
