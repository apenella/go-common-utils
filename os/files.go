package common

import (
	"io/ioutil"
	"os"
)

// FileExist return an error when a file does not exist
func FileExist(file string) error {
	_, err := os.Stat(file)
	return err
}

// GetFolderFiles get all files from folder
func GetFolderFiles(folder string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(folder)

	if err != nil {
		return nil, err
	}

	return files, nil
}
