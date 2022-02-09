package archiver

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// listFolder returns a slice with all root folder file paths
func listFolder(root string) []string {
	var fls []string

	folderInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("Error listing file paths: root folder not found")
	}

	for _, fileInfo := range folderInfo {
		if !fileInfo.IsDir() {
			fls = append(fls, filepath.Join(root, fileInfo.Name()))
		} else {
			fls = append(fls, listFolder(filepath.Join(root, fileInfo.Name()))...)
		}
	}
	return fls
}

func readFile(fileName string) ([]byte, error) {
	var err error

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
