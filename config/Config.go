package config

import (
	"os"
	"path"
	"path/filepath"
)

func getConfigFileFromExecutable(fileName string) *os.File {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return  nil
	}

	f, err := os.Open(path.Join(dir, fileName))

	if err != nil {
		return nil
	}

	return f
}