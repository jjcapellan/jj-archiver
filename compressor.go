package archiver

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Unzip takes a gzip file (src) and uncompress it in dst directory.
//
// If destination is not defined (dst == "") then uses current directory.
func UnZip(src string, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}

	zr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	fName := zr.Name

	defer zr.Close()
	defer f.Close()

	result, _ := ioutil.ReadAll(zr)

	if dst == "" {
		dst = fName
	} else {
		if _, err := os.Stat(dst); os.IsNotExist(err) {
			os.MkdirAll(dst, 0777)
		}
		dst = filepath.Join(dst, fName)
	}

	err = os.WriteFile(dst, result, 0777)
	if err != nil {
		return err
	}

	err = zr.Close()
	if err != nil {
		return err
	}

	return nil

}

// Zip takes a file (src) a compress it using gzip algorithm into
// a defined directory (dstDir) with ".gz" file extension.
//
// If dstDir == "" then gz file is saved in current directory.
//
// Example: Zip("folder1/file.tar", "") produces "./file.tar.gz"
func Zip(src string, dstDir string) error {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)

	b, err := readFile(src)
	if err != nil {
		return err
	}

	// Saves original name in header
	_, zw.Header.Name = filepath.Split(src)

	_, err = zw.Write(b)
	if err != nil {
		return err
	}

	err = zw.Close()
	if err != nil {
		return err
	}

	_, fName := filepath.Split(src)
	if dstDir == "" {
		dstDir = fName
	} else {
		dstDir = filepath.Join(dstDir)
		if _, err := os.Stat(dstDir); os.IsNotExist(err) {
			os.MkdirAll(dstDir, 0777)
		}
		dstDir = filepath.Join(dstDir, fName)
	}
	dstDir += ".gz"

	err = os.WriteFile(dstDir, buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
