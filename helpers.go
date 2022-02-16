package archiver

import (
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func GetCRC32(fileName string) uint32 {
	b, _ := readFile(fileName)
	crc := crc32.Checksum(b, crc32.IEEETable)
	return crc
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

func prepareDst(src string, dstDir string, ext string, removeExt bool) (dst string, e error) {

	_, fName := filepath.Split(src)
	if removeExt {
		fName = fName[:(len(fName) - len(ext))]
	}
	if dstDir == "" {
		dstDir = fName
	} else {
		dstDir = filepath.Join(dstDir)
		if _, err := os.Stat(dstDir); os.IsNotExist(err) {
			err = os.MkdirAll(dstDir, 0777)
			if err != nil {
				return "", err
			}
		}
		dstDir = filepath.Join(dstDir, fName)
	}
	if !removeExt {
		dstDir += ext
	}
	return dstDir, nil
}

func readFile(fileName string) ([]byte, error) {
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
