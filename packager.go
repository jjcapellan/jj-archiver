package archiver

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func writeTarHeader(path string, tw *tar.Writer) error {
	fInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	fHeader, err := tar.FileInfoHeader(fInfo, "")
	if err != nil {
		return err
	}

	fHeader.Name = path
	err = tw.WriteHeader(fHeader)
	if err != nil {
		return err
	}

	return nil
}

func writeTarBody(path string, tw *tar.Writer) error {
	body, err := ReadFile(path)
	if err != nil {
		return err
	}
	_, err = tw.Write(body)
	if err != nil {
		return err
	}
	return nil
}

// PackFolder packs input folder into an array of bytes.
// This array can be used to write a tar file, or to process in another function.
//
// Example: PackFolder("user/projectsfolder") returns []byte of the packed file
func PackFolder(input string) ([]byte, error) {
	var buffer bytes.Buffer

	tw := tar.NewWriter(&buffer)

	basePath := filepath.Dir(input)

	fileNames, dirNames := listFolder(input, basePath)

	for _, path := range dirNames {
		err := writeTarHeader(path, tw)
		if err != nil {
			return nil, err
		}
	}

	for _, path := range fileNames {
		err := writeTarHeader(path, tw)
		if err != nil {
			return nil, err
		}
		err = writeTarBody(path, tw)
		if err != nil {
			return nil, err
		}
	}

	tw.Close()

	data := buffer.Bytes()

	return data, nil
}

// Unpack extracts all files []byte of the input tar file to output path.
//
// Example: Unpack(fileBytesArray, "unpackedfolders/folder1")
func Unpack(input []byte, output string) error {

	file := bytes.NewReader(input)

	tr := tar.NewReader(file)

	for {
		fHeader, err := tr.Next()

		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		path := filepath.Join(output, fHeader.Name)

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
