package consensus

import "os"

func ensureDirectory(path string, mode os.FileMode) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(path, mode | os.ModeDir)
	}
}