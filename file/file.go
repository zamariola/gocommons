package file

import (
	"io"
	"os"
)

//Copy function is responsible for copying the file contents from a
//source path to a destination path. It overwrites existing destination files
//and does not copy attributes
func Copy(source, destination string) error {
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}
	return output.Close()
}

//Exists verify if a file or directory is present
//in the file system
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
