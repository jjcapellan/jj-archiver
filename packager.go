package archiver

import (
	"archive/tar"
	"bytes"
	"os"
)

// PackFolder packages a folder into a tar file ("output.tar")
func PackFolder(folder string, output string) error {
	var buffer bytes.Buffer

	tw := tar.NewWriter(&buffer)

	filePaths := listFolder(folder)

	for _, path := range filePaths {
		var err error
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		// File header
		fInfo, err := f.Stat()
		if err != nil {
			return err
		}
		fHeader, err := tar.FileInfoHeader(fInfo, "")
		if err != nil {
			return err
		}
		fHeader.Name = path

		// File body
		body, err := readFile(path)
		if err != nil {
			return err
		}

		// Write header
		err = tw.WriteHeader(fHeader)
		if err != nil {
			return err
		}

		// Write body
		_, err = tw.Write(body)
		if err != nil {
			return err
		}

		// Close current file
		err = f.Close()
		if err != nil {
			return err
		}
	} //End filePaths range

	tw.Close()

	err := os.WriteFile(output+".tar", buffer.Bytes(), 0666)
	if err != nil {
		return err
	}

	return nil
}
