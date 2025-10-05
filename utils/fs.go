package utils

import (
	"os"
	"path/filepath"
)

type Fs struct{}

func CreateFs() *Fs {
	return &Fs{}
}

func (fs Fs) Create(name string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if filepath.IsAbs(name) {
		return createFile(name)
	} else {
		return createFile(filepath.Join(dir, name))
	}
}

func fileExists(path string) bool {
	stat, err := os.Stat(path)

	if err != nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true
}

func folderExists(name string) bool {
	info, err := os.Stat(name)

	if err != nil {
		return false
	}

	if !info.IsDir() {
		return false
	}

	return true
}

func createFile(name string) error {
	if fileExists(name) {
		return nil
	}

	if err := createFolder(filepath.Dir(name)); err != nil {
		return err
	}

	_, err := os.Create(name)

	return err
}

func createFolder(name string) error {
	if folderExists(name) {
		return nil
	}

	return os.MkdirAll(name, 0755)
}
