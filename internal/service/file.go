package service

import (
	"os"
	"path"
)

// FileService implements the methods for file operations
type FileService struct {
}

// Rename renames the file with the specified oldPath and newName.
// It returns the new path after the renaming operation.
func (fs FileService) Rename(oldPath, newName string) (string, error) {
	if newName == "" {
		return oldPath, nil
	}
	newPath := path.Join(path.Dir(oldPath), newName)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return oldPath, err
	}
	return newPath, nil
}

// NewFileService creates a new instance of FileService.
func NewFileService() *FileService {
	return &FileService{}
}
