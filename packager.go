package archiver

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
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

func Unpack(src string, dst string) error {

	file, err := os.Open(src)
	defer file.Close()
	if err != nil {
		return err
	}

	tr := tar.NewReader(file)

	for {
		var err error
		fHeader, err := tr.Next()

		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		path := filepath.Join(dst, fHeader.Name)

		if fHeader.Typeflag == tar.TypeDir {
			os.MkdirAll(path, os.FileMode(fHeader.Mode))
		}

		if fHeader.Typeflag == tar.TypeReg {

			buffer := make([]byte, fHeader.Size)
			tr.Read(buffer)

			f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.FileMode(fHeader.Mode))
			if err != nil {
				return err
			}

			_, err = f.Write(buffer)
			if err != nil {
				return err
			}

			err = f.Close()
			if err != nil {
				return err
			}
		}
	} // End for
}
