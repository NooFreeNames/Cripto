// Package service implements the service layer of the application.
package service

// FileCipherInput represents the input parameters for file encryption
// or decryption.
type FileCipherInput struct {
	Path     string // The path to the file.
	Password string // The password used for encryption or decryption.
}

// FileCipher defines the methods for file encryption and decryption.
type FileCipher interface {
	// Encrypt encrypts the file using the provided input parameters.
	Encrypt(input FileCipherInput) error
	// Decrypt decrypts the file using the provided input parameters.
	Decrypt(input FileCipherInput) error
}

// File defines the methods for file operations.
type File interface {
	// Rename renames the file with the specified old path and new name.
	Rename(oldPath string, newName string) (string, error)
}

// Service represents application services
type Service struct {
	FileCipher
	File
}

// NewService creates a new instance of the Service.
func NewService() *Service {
	return &Service{
		FileCipher: NewFileCipherService(),
		File:       NewFileService(),
	}
}
