package archiver

import (
	"archive/tar"
	"bytes"
	"os"
)

func writeTarHeader(path string, tw *tar.Writer) {
	fInfo, _ := os.Stat(path)
	fHeader, _ := tar.FileInfoHeader(fInfo, "")
	fHeader.Name = path
	tw.WriteHeader(fHeader)
}

func writeTarBody(path string, tw *tar.Writer) {
	body, _ := readFile(path)
	tw.Write(body)
}

// PackFolder packages a folder into a tar file ("output.tar")
func PackFolder(folder string, output string) error {
	var buffer bytes.Buffer

	tw := tar.NewWriter(&buffer)

	fileNames, dirNames := listFolder(folder)

	for _, path := range dirNames {
		writeTarHeader(path, tw)
	}

	for _, path := range fileNames {
		writeTarHeader(path, tw)
		writeTarBody(path, tw)
	}

	tw.Close()

	err := os.WriteFile(output+".tar", buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
