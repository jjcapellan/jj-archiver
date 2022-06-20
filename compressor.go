package archiver

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Decompress decompress bytes array of file into another bytes array.
func Decompress(input []byte) (output []byte, fileName string, e error) {

	f := bytes.NewReader(input)

	zr, err := gzip.NewReader(f)
	if err != nil {
		return nil, "", err
	}

	fName := zr.Name

	defer zr.Close()

	result, _ := ioutil.ReadAll(zr)

	err = zr.Close()
	if err != nil {
		return nil, "", err
	}

	return result, fName, nil

}

// Compress compress bytes array of file into another bytes array.
//
// fileName could include relative file path. It is stored in the header of the compressed file.
//
// This param is important to preserve the original name of the file when uncompressed.
//
// Example: Compress(filedata, "folder/uncompressedfile.ext")
func Compress(input []byte, fileName string) ([]byte, error) {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)

	// Saves original name in header
	_, zw.Header.Name = filepath.Split(fileName)

	_, err := zw.Write(input)
	if err != nil {
		return nil, err
	}

	err = zw.Close()
	if err != nil {
		return nil, err
	}

	data := buffer.Bytes()

	return data, nil
}

// GetGzipCRC32 gets crc32 checksum of decompressed file stored in gzip file (fileName)
func GetGzipCRC32(fileName string) (uint32, error) {
	var size uint32
	footer, err := getGzipFooter(fileName)
	if err != nil {
		return 0, err
	}

	if isLittleEndian() {
		binary.LittleEndian.Uint32(footer[:4])
	} else {
		binary.BigEndian.Uint32(footer[:4])
	}

	return size, nil
}

// GetDecompressedSize gets decompresed size of the gzip file.
func GetDecompressedSize(fileName string) (uint32, error) {
	var size uint32
	footer, err := getGzipFooter(fileName)
	if err != nil {
		return 0, err
	}

	if isLittleEndian() {
		size = binary.LittleEndian.Uint32(footer[4:])
	} else {
		size = binary.BigEndian.Uint32(footer[4:])
	}
	return size, nil
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
