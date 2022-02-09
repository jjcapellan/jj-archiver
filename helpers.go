package archiver

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

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
