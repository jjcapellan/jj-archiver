package archiver

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Decompress decompress input file into output path.
//
// If output == "" then uses current directory.
//
// Example: Decompress("projects.gz", "user/projectsfolder")
func Decompress(input string, output string) error {
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

// Compress compress input file into output path. ".gz" extension is added to output.
//
// If output == "" then generated file is saved in current directory.
//
// Example: Compress("folder1/myfile.tar", "") generates "./myfile.tar.gz"
func Compress(input string, output string) error {
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

	dst, err := prepareDst(input, output, extCompressed, false)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}

// GetDecompressedSize gets decompresed size of the gzip file.
func GetDecompressedSize(fileName string) (uint32, error) {
	footer, err := getGzipFooter(fileName)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(footer[4:]), nil
}

func getGzipFooter(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	// Last 8 bytes of file -> CRC32 checksum of uncompresed file (4 bytes) + uncompresed file size (4 bytes)
	buffer := make([]byte, 8)

	_, err = f.ReadAt(buffer, fileSize-8)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
