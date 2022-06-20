package archiver

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Encrypt encrypts input []byte into output []byte using the password.
// This function uses AES256 algorithm (mode GCM).
//
// Password lenght can be any not zero value. The password is processed by
// the SHA256 hash algorithm to generate a 256-bit key.
func Encrypt(input []byte, password string) (output []byte, e error) {

	gcm, err := getGCM(password)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	result := gcm.Seal(nonce, nonce, input, nil)

	return result, nil
}

// Decrypt decrypts input []byte] into output []byte using the password.
// This function uses AES256 algorithm (mode GCM).
func Decrypt(input []byte, password string) (output []byte, e error) {

	gcm, err := getGCM(password)
	if err != nil {
		return nil, err
	}

	s := gcm.NonceSize()

	nonce, cipherContent := input[:s], input[s:]

	result, err := gcm.Open(nil, nonce, cipherContent, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
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
