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
func UnZip(input string, output string) error {
	f, err := os.Open(input)
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

	if output == "" {
		output = fName
	} else {
		if _, err := os.Stat(output); os.IsNotExist(err) {
			os.MkdirAll(output, 0777)
		}
		output = filepath.Join(output, fName)
	}

	err = os.WriteFile(output, result, 0777)
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
func Zip(input string, output string) error {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)

	b, err := readFile(input)
	if err != nil {
		return err
	}

	// Saves original name in header
	_, zw.Header.Name = filepath.Split(input)

	_, err = zw.Write(b)
	if err != nil {
		return err
	}

	err = zw.Close()
	if err != nil {
		return err
	}

	dst, err := prepareDst(input, output, ".gz", false)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
