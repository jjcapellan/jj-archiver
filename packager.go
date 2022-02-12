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
	body, err := readFile(path)
	if err != nil {
		return err
	}
	_, err = tw.Write(body)
	if err != nil {
		return err
	}
	return nil
}

// PackFolder packs input folder into output tar file.
// ".tar" extension is added to the output file name.
//
// Example: PackFolder("user/projectsfolder", "projects") generates "./projects.tar"
func PackFolder(input string, output string) error {
	var buffer bytes.Buffer

	tw := tar.NewWriter(&buffer)

	fileNames, dirNames := listFolder(input)

	for _, path := range dirNames {
		err := writeTarHeader(path, tw)
		if err != nil {
			return err
		}
	}

	for _, path := range fileNames {
		err := writeTarHeader(path, tw)
		if err != nil {
			return err
		}
		err = writeTarBody(path, tw)
		if err != nil {
			return err
		}
	}

	tw.Close()

	dir, fName := filepath.Split(output)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	err := os.WriteFile(filepath.Join(dir, fName)+extPacked, buffer.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}

// Unpack extracts all files from input tar file to output path.
//
// Example: Unpack("packedfiles/myfile.tar", "unpackedfolders/folder1")
func Unpack(input string, output string) error {

	file, err := os.Open(input)
	defer file.Close()
	if err != nil {
		return err
	}

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
