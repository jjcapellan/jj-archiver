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

// Compress compress bytes array of file into another bytes array.
//
// fileName could include relative file path. It is stored in the header of the compressed file.
//
// This param is important to preserve the original name of the file when uncompressed.
//
// Example: Compress(filedata, "folder/uncompressedfile.ext")
func Compress(fileData []byte, fileName string) (error, []byte) {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)

	// Saves original name in header
	_, zw.Header.Name = filepath.Split(fileName)

	_, err := zw.Write(fileData)
	if err != nil {
		return err, nil
	}

	err = zw.Close()
	if err != nil {
		return err, nil
	}

	data := buffer.Bytes()

	return nil, data
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
