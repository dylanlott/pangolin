package utils

import "os"

func EnsureDirectory(path string, mode os.FileMode) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(path, mode)
	}

	return err
}
