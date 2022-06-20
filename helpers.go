package archiver

import (
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

// GetCRC32 gets the crc32 of any file using IEEE polynomial
func GetCRC32(fileName string) (uint32, error) {
	b, err := ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	crc := crc32.Checksum(b, crc32.IEEETable)
	return crc, nil
}

func isLittleEndian() bool {
	p1 := new(int16)
	*p1 = 1 // littleendian: [0x01 0x00] bigEndian: [0x00 0x01]
	p2 := (*int8)(unsafe.Pointer(p1))
	return *p2 == 1
}

// listFolder returns 2 slices with files and folders paths
func listFolder(root string) (files []string, folders []string) {
	var fls []string
	var dirs []string

	folderInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("Error listing file paths: root folder not found")
	}

	for _, fileInfo := range folderInfo {
		if !fileInfo.IsDir() {
			fls = append(fls, filepath.Join(root, fileInfo.Name()))
		} else {
			dirs = append(dirs, filepath.Join(root, fileInfo.Name()))
			subFls, subDirs := listFolder(filepath.Join(root, fileInfo.Name()))
			dirs = append(dirs, subDirs...)
			fls = append(fls, subFls...)
		}
	}
	return fls, dirs
}

func ReadFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	size := fileInfo.Size()

	buffer := make([]byte, size)

	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	dir, fileName := filepath.Split(filePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	err := os.WriteFile(filepath.Join(dir, fileName), data, perm)

	return err
}
