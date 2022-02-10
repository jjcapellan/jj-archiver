package archiver

import (
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
)

// Zip takes a file (src) a compress it using gzip algorithm into
// a defined place (dst) with ".gz" file extension.
//
// If dst == "" then gz file is saved in current directory.
//
// Example: Zip("folder1/file.tar", "") produces "./file.tar.gz"
func Zip(src string, dst string) error {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)

	b, err := readFile(src)
	if err != nil {
		return err
	}

	_, err = zw.Write(b)
	if err != nil {
		return err
	}

	err = zw.Close()
	if err != nil {
		return err
	}

	dir, fName := filepath.Split(dst)
	if dst == "" {
		_, fName = filepath.Split(src)
		dst = fName + ".gz"
	} else {
		dst = filepath.Join(dir, fName) + ".gz"
	}
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0777)
		}
	}

	err = os.WriteFile(dst, buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
